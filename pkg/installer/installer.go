package installer

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 用于为本机安装maven

// Install 根据当前操作系统选择合适的安装方法
func Install() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return InstallWindows()
	case "linux":
		return InstallLinux()
	case "darwin":
		return InstallMacOS()
	default:
		return "", fmt.Errorf("暂不支持当前操作系统: %s", runtime.GOOS)
	}
}

// 下载文件（通用函数）
func downloadFile(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，HTTP状态码: %d", resp.StatusCode)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

// 解压tar.gz文件
func untar(tarPath, destDir string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	// 创建目标目录
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(destDir, header.Name)

		// 检查路径穿越漏洞
		if !strings.HasPrefix(path, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("非法的文件路径: %s", path)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}

			outFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	return nil
}

// 从tar.gz包安装Maven的通用函数
func installFromTarGz(mavenURL string) (string, error) {
	// 安装目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("获取用户主目录失败: %w", err)
	}

	// 创建安装目录
	mavenDir := filepath.Join(homeDir, ".m2", "maven")
	if err := os.MkdirAll(mavenDir, 0755); err != nil {
		return "", fmt.Errorf("创建Maven安装目录失败: %w", err)
	}

	// 下载Maven
	tarPath := filepath.Join(mavenDir, "maven.tar.gz")
	if err := downloadFile(mavenURL, tarPath); err != nil {
		return "", fmt.Errorf("下载Maven失败: %w", err)
	}

	// 解压Maven
	extractDir := filepath.Join(mavenDir, "maven-install")
	if err := untar(tarPath, extractDir); err != nil {
		return "", fmt.Errorf("解压Maven失败: %w", err)
	}

	// 查找解压后的目录
	mavenHome := ""
	err = filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.Contains(info.Name(), "apache-maven") {
			mavenHome = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("查找Maven目录失败: %w", err)
	}

	if mavenHome == "" {
		return "", errors.New("找不到Maven安装目录")
	}

	// 确保bin目录存在
	binDir := filepath.Join(mavenHome, "bin")
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		return "", errors.New("Maven安装不完整，找不到bin目录")
	}

	// 删除临时文件
	os.Remove(tarPath)

	return mavenHome, nil
}
