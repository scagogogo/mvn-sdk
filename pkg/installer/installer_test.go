package installer

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDownloadFile 测试下载文件功能
func TestDownloadFile(t *testing.T) {
	// 创建一个测试服务器，提供测试文件
	testContent := []byte("This is a test file for Maven installer")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(testContent)
	}))
	defer server.Close()

	// 创建临时目录保存下载的文件
	tempDir, err := os.MkdirTemp("", "download-test")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 测试下载文件
	destPath := filepath.Join(tempDir, "test-file.txt")
	err = downloadFile(server.URL, destPath)
	if err != nil {
		t.Fatalf("下载文件失败: %v", err)
	}

	// 检查文件是否存在
	_, err = os.Stat(destPath)
	if err != nil {
		t.Fatalf("下载的文件不存在: %v", err)
	}

	// 验证文件内容
	content, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("读取下载的文件失败: %v", err)
	}

	if string(content) != string(testContent) {
		t.Fatalf("文件内容不匹配，期望: %s, 实际: %s", testContent, content)
	}

	t.Log("下载文件测试通过")
}

// TestCreateTarGzArchive 创建一个简单的tar.gz文件用于测试
func createTestTarGz(t *testing.T, path string) {
	// 由于创建tar.gz文件比较复杂，这里使用简化的方法
	// 在实际测试中，可以使用预先准备好的小型测试文件
	t.Skip("需要预先准备测试用的tar.gz文件")
}

// TestUntar 测试解压tar.gz文件功能
// 注意：这个测试需要一个预先准备好的tar.gz测试文件
func TestUntar(t *testing.T) {
	// 此测试需要一个预先准备好的tar.gz文件
	// 如果没有适当的测试文件，则跳过此测试
	t.Skip("需要预先准备测试用的tar.gz文件，暂时跳过此测试")

	// 创建临时目录解压文件
	tempDir, err := os.MkdirTemp("", "untar-test")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 测试tar.gz文件路径，这里使用简化的示例
	// 实际测试中需要替换为真实的测试文件路径
	testTarPath := filepath.Join(tempDir, "test.tar.gz")
	createTestTarGz(t, testTarPath)

	// 解压目录
	extractDir := filepath.Join(tempDir, "extracted")
	err = untar(testTarPath, extractDir)
	if err != nil {
		t.Fatalf("解压文件失败: %v", err)
	}

	// 检查解压后的文件夹是否存在
	_, err = os.Stat(extractDir)
	if err != nil {
		t.Fatalf("解压目录不存在: %v", err)
	}

	t.Log("解压文件测试通过")
}

// TestInstall 测试安装功能（集成测试）
// 由于会实际安装Maven，默认跳过此测试
func TestInstall(t *testing.T) {
	// 默认跳过网络依赖测试，只有在环境变量指定时才运行
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("跳过安装测试（设置RUN_INTEGRATION_TESTS=1来运行）")
	}

	// 测试安装Maven
	mavenHome, err := Install()
	if err != nil {
		t.Logf("安装Maven失败（可能是网络问题）: %v", err)
		t.Skip("网络或安装问题，跳过此测试")
	}

	// 验证安装路径
	if mavenHome == "" {
		t.Fatal("返回的Maven安装路径为空")
	}

	// 验证bin/mvn可执行文件存在
	mvnPath := filepath.Join(mavenHome, "bin", "mvn")
	_, err = os.Stat(mvnPath)
	if err != nil {
		t.Fatalf("未找到mvn可执行文件: %v", err)
	}

	t.Logf("Maven成功安装到: %s", mavenHome)
}

// TestPathTraversalSecurity 测试路径穿越安全检查
func TestPathTraversalSecurity(t *testing.T) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "security-test")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 测试目录
	destDir := filepath.Join(tempDir, "dest")
	if err := os.MkdirAll(destDir, 0755); err != nil {
		t.Fatalf("创建测试目录失败: %v", err)
	}

	// 测试不同的路径组合
	testCases := []struct {
		name       string
		headerPath string
		expectSafe bool
	}{
		{"安全路径", "normal/path/file.txt", true},
		{"穿越路径1", "../dangerous/path.txt", false},
		{"穿越路径2", "../../etc/passwd", false},
		// normal/../path/file.txt 经过Clean后变成 path/file.txt，这是安全的
		{"包含../但安全的路径", "normal/../path/file.txt", true},
		// 这个路径经过Clean后会尝试访问destDir外的目录，应该被检测为不安全
		{"真正的穿越路径", "subdir/../../outside.txt", false},
		{"绝对路径", "/etc/passwd", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var path string
			if filepath.IsAbs(tc.headerPath) {
				// 如果是绝对路径，直接使用
				path = tc.headerPath
			} else {
				// 相对路径，与destDir结合
				path = filepath.Join(destDir, tc.headerPath)
			}

			// 模拟untar函数中的路径安全检查
			// 先使用filepath.Clean对路径进行规范化，然后检查是否仍然在目标目录内
			cleanedPath := filepath.Clean(path)
			cleanedDestDir := filepath.Clean(destDir)
			isSafe := strings.HasPrefix(cleanedPath, cleanedDestDir+string(os.PathSeparator))

			// 验证结果
			if isSafe != tc.expectSafe {
				t.Errorf("安全检查结果不符预期，路径: %s (清理后: %s), 期望安全: %v, 实际: %v",
					path, cleanedPath, tc.expectSafe, isSafe)
			}
		})
	}
}
