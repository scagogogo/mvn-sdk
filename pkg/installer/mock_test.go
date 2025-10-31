package installer

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// 测试用的小型Maven结构
type mockMaven struct {
	Version   string
	BinFiles  []string
	LibFiles  []string
	ConfFiles []string
}

// 创建一个模拟的Maven结构用于测试
func createMockMaven() mockMaven {
	return mockMaven{
		Version:  "3.9.11",
		BinFiles: []string{"mvn", "mvnDebug"},
		LibFiles: []string{"maven-core.jar", "maven-model.jar"},
		ConfFiles: []string{
			"settings.xml",
			"logging/simplelogger.properties",
		},
	}
}

// 创建一个测试用的Maven tar.gz文件
func createMockMavenTarGz(t *testing.T, outputPath string) error {
	mock := createMockMaven()

	// 创建目录结构
	baseName := "apache-maven-" + mock.Version

	// 创建缓冲区用于写入tar.gz文件
	var buf bytes.Buffer
	gzWriter := gzip.NewWriter(&buf)
	tarWriter := tar.NewWriter(gzWriter)

	// 添加根目录
	header := &tar.Header{
		Name:     baseName + "/",
		Mode:     0755,
		Typeflag: tar.TypeDir,
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	// 添加bin目录
	header = &tar.Header{
		Name:     baseName + "/bin/",
		Mode:     0755,
		Typeflag: tar.TypeDir,
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	// 添加bin下的可执行文件
	for _, file := range mock.BinFiles {
		content := []byte("#!/bin/sh\necho This is a mock Maven binary")
		header = &tar.Header{
			Name:     baseName + "/bin/" + file,
			Mode:     0755,
			Size:     int64(len(content)),
			Typeflag: tar.TypeReg,
		}
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}
		if _, err := tarWriter.Write(content); err != nil {
			return err
		}
	}

	// 添加lib目录
	header = &tar.Header{
		Name:     baseName + "/lib/",
		Mode:     0755,
		Typeflag: tar.TypeDir,
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	// 添加lib下的jar文件
	for _, file := range mock.LibFiles {
		content := []byte("Mock jar file content")
		header = &tar.Header{
			Name:     baseName + "/lib/" + file,
			Mode:     0644,
			Size:     int64(len(content)),
			Typeflag: tar.TypeReg,
		}
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}
		if _, err := tarWriter.Write(content); err != nil {
			return err
		}
	}

	// 添加conf目录
	header = &tar.Header{
		Name:     baseName + "/conf/",
		Mode:     0755,
		Typeflag: tar.TypeDir,
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	// 添加conf文件
	for _, file := range mock.ConfFiles {
		dir := filepath.Dir(file)
		if dir != "." {
			// 为嵌套目录创建目录条目
			header = &tar.Header{
				Name:     baseName + "/conf/" + dir + "/",
				Mode:     0755,
				Typeflag: tar.TypeDir,
			}
			if err := tarWriter.WriteHeader(header); err != nil {
				return err
			}
		}

		content := []byte("# Mock configuration file\n")
		header = &tar.Header{
			Name:     baseName + "/conf/" + file,
			Mode:     0644,
			Size:     int64(len(content)),
			Typeflag: tar.TypeReg,
		}
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}
		if _, err := tarWriter.Write(content); err != nil {
			return err
		}
	}

	// 关闭writer
	if err := tarWriter.Close(); err != nil {
		return err
	}
	if err := gzWriter.Close(); err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}

// 创建测试用的模拟HTTP服务器，提供Maven下载
func createMockMavenServer(t *testing.T) (*httptest.Server, string) {
	mux := http.NewServeMux()

	// 创建临时目录用于存储模拟文件
	tempDir, err := os.MkdirTemp("", "maven-mock")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}

	// 创建模拟的Maven tar.gz文件
	mockTarPath := filepath.Join(tempDir, "apache-maven-3.9.11-bin.tar.gz")
	if err := createMockMavenTarGz(t, mockTarPath); err != nil {
		t.Fatalf("创建模拟Maven包失败: %v", err)
	}

	// 处理Maven下载请求
	mux.HandleFunc("/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(mockTarPath)
		if err != nil {
			http.Error(w, "读取文件失败", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=apache-maven-3.9.11-bin.tar.gz")
		w.Write(data)
	})

	server := httptest.NewServer(mux)
	return server, tempDir
}

// 使用模拟服务器和可测试版本的函数进行测试
func TestInstallMacOSWithMock(t *testing.T) {
	// 创建模拟HTTP服务器
	server, tempDir := createMockMavenServer(t)
	defer os.RemoveAll(tempDir)
	defer server.Close()

	// 设置模拟URL
	mockURL := server.URL + "/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.tar.gz"

	// 创建测试目录
	testHomeDir, err := os.MkdirTemp("", "maven-test-home")
	if err != nil {
		t.Fatalf("创建测试主目录失败: %v", err)
	}
	defer os.RemoveAll(testHomeDir)

	// 设置测试选项
	options := InstallOptions{
		MavenURL:     mockURL,
		HomeDir:      testHomeDir,
		SkipEnvSetup: true, // 跳过环境变量设置
	}

	// 使用可测试版本的函数
	mavenHome, err := InstallMacOSWithOptions(options)
	if err != nil {
		t.Fatalf("使用模拟服务器安装Maven失败: %v", err)
	}

	// 验证安装结果
	if mavenHome == "" {
		t.Fatal("返回的Maven安装路径为空")
	}

	// 验证bin/mvn可执行文件存在
	mvnPath := filepath.Join(mavenHome, "bin", "mvn")
	_, err = os.Stat(mvnPath)
	if err != nil {
		t.Fatalf("未找到mvn可执行文件: %v", err)
	}

	t.Logf("Maven成功安装到测试目录: %s", mavenHome)
}
