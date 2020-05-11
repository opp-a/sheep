# sheep

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

### Lints and fixes files
```
yarn lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

vscode设置：
{
    "editor.formatOnSave": true,
    "editor.tabSize": 2,
    "editor.quickSuggestions": {
        "other": true,
        "comments": true,
        "strings": true
    },
    "editor.codeActionsOnSave": {
        "source.fixAll.eslint": true
    },
    "files.associations": {
        "*.vue": "html"
    },
    "eslint.validate": [
        "javascript",
        "javascriptreact",
        {
            "language": "vue",
            "autoFix": true
        }
    ],
    "htmlhint.options": {
        "tagname-lowercase": false,
        "attr-lowercase": true,
        "attr-value-double-quotes": true,
        "doctype-first": false
    },
    "prettier.trailingComma": "none",
    "prettier.proseWrap": "never",
    "prettier.printWidth": 120,
    "prettier.tabWidth": 2,
    "prettier.jsxSingleQuote": true,
    "prettier.singleQuote": true,
    "prettier.semi": false,
    "editor.suggest.snippetsPreventQuickSuggestions": false,
    "explorer.confirmDragAndDrop": false,
    "prettier.htmlWhitespaceSensitivity": "strict",
    "prettier.bracketSpacing": false
}
