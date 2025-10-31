module.exports = {
  title: 'Maven SDK Go',
  description: '用于 Maven 操作的 Go SDK',
  
  themeConfig: {
    nav: [
      { text: '首页', link: '/' },
      { text: 'API 参考', link: '/api' }
    ],

    sidebar: [
      {
        text: '指南',
        items: [
          { text: '快速开始', link: '/' },
          { text: 'API 参考', link: '/api' }
        ]
      },
      {
        text: '包',
        items: [
          { text: 'Finder', link: '/api#finder' },
          { text: 'Command', link: '/api#command' },
          { text: 'Local Repository', link: '/api#local-repository' },
          { text: 'Installer', link: '/api#installer' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/scagogogo/mvn-sdk' }
    ],

    footer: {
      message: '基于 MIT 许可证发布。',
      copyright: '版权所有 © 2024 scagogogo'
    }
  }
}