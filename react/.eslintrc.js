// module.exports = {
//     "parser": "@babel/eslint-parser",
//     "rules": {}
// }
module.exports = {
    "env": {
        "browser": true
    },
    "plugins": [ "react" ],
    "extends": [
      "eslint:recommended",
      "plugin:react/recommended"
    ],
    "parser": "@babel/eslint-parser",
    "workingDirectories": [
        {"mode": "auto"}   // Path to the shared code directory
      ],
};