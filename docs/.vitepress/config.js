module.exports = {
  title: 'Maven SDK Go',
  description: 'A Go SDK for Maven operations',
  base: process.env.NODE_ENV === 'production' ? '/mvn-sdk/' : '/',
  
  // 多语言配置
  locales: {
    root: {
      label: 'English',
      lang: 'en-US'
    },
    zh: {
      label: '简体中文',
      lang: 'zh-CN'
    }
  },
  
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'API Reference', link: '/api' }
    ],

    sidebar: [
      {
        text: 'Guide',
        items: [
          { text: 'Getting Started', link: '/' },
          { text: 'API Reference', link: '/api' }
        ]
      },
      {
        text: 'Packages',
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
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024 scagogogo'
    }
  }
}