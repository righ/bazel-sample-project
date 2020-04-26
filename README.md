This is a bazel sample project.

# requirements

- bazel
- skaffold

# How to use

- `make skaffold`
- go `http://localhost`.
- exec the following code on the site:
  ```js
  fetch("/echo", {method: "post", body: "aaa"}).then(res => res.text()).then(text => console.log(text))
  // aaa
  ```
