module.exports = {
  env: {
    es2021: true,
    node: true,
  },
  extends: "airbnb-base",
  overrides: [
  ],
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module",
  },
  rules: {
    "import/extensions": ["error", "ignorePackages"],
    "no-unused-vars": ["error", {
      vars: "all", args: "after-used", argsIgnorePattern: "^_", ignoreRestSiblings: false,
    }],
    quotes: ["error", "double"],
    "no-console": ["error", { allow: ["error"] }],
    "no-use-before-define": [
      "error", {
        functions: false,
        classes: true,
        variables: true,
      },
    ],
  },
};
