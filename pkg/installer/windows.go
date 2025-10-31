package installer

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallWindows 在Windows系统上安装Maven
func InstallWindows() (string, error) {
	// Maven下载地址
	mavenURL := "https://archive.apache.org/dist/maven/maven-3/3.9.11/binaries/apache-maven-3.9.11-bin.zip"
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
	zipPath := filepath.Join(mavenDir, "maven.zip")
	if err := downloadFileWindows(mavenURL, zipPath); err != nil {
		return "", fmt.Errorf("下载Maven失败: %w", err)
	}

	// 解压Maven
	extractDir := filepath.Join(mavenDir, "maven-install")
	if err := unzipWindows(zipPath, extractDir); err != nil {
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

	// 设置环境变量
	if err := setEnvVarsWindows(mavenHome); err != nil {
		return "", fmt.Errorf("设置环境变量失败: %w", err)
	}

	// 删除临时zip文件
	os.Remove(zipPath)

	return mavenHome, nil
}

// 下载文件
func downloadFileWindows(url, destPath string) error {
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

// 解压zip文件
func unzipWindows(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	// 创建目标目录
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)

		// 检查路径穿越漏洞
		if !strings.HasPrefix(fpath, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("非法的文件路径: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

// 设置环境变量
func setEnvVarsWindows(mavenHome string) error {
	// Windows上使用SETX命令持久化环境变量
	cmds := []struct {
		name string
		args []string
	}{
		{"setx", []string{"MAVEN_HOME", mavenHome}},
		{"setx", []string{"PATH", fmt.Sprintf("%%PATH%%;%s", filepath.Join(mavenHome, "bin"))}},
	}

	for _, cmd := range cmds {
		c := exec.Command(cmd.name, cmd.args...)
		if err := c.Run(); err != nil {
			return fmt.Errorf("执行命令 %s %v 失败: %w", cmd.name, cmd.args, err)
		}
	}

	return nil
}
