# Grator 🚀

Grator is a versatile load testing library designed for assessing the performance and scalability of the entire flow of your application.
This version of Grator is implemented in Golang and empowers you to simulate real-world scenarios and evaluate system behavior under multi user complete flow load.

## Features 🌟

- **Flow Testing**: Target sequential processes and test end to end performance under varying loads.

- **Concurrency**: Define custom load to simulate different levels of concurrent users or transactions.

- **Templating**: Store values from responses into session values and template them into subsequent requests.

- **Real-time Monitoring**: *Work In Progress*

- **Extensive Reporting**: *Work In Progress*

## Getting Started 🚦

### Installation

> ```bash
> go get -u github.com/digi0ps/grator
> ```

### Example Usage

Create a YAML file, e.g., `config.yml`:

```yaml
baseURL: "http://<baseURL>"
maxActors: 1
actions:
- method: GET
  url: /api/ping/

- method: POST
  url: /api/otp/generate
  body: "{\"phone_number\": \"9876543210\"}"
  storeValues:
    got_otp: data.otp

- method: POST
  url: /api/otp/customer/authenticate
  body: "{\"phone_number\": \"9876543210\", \"otp\": \"%got_otp%\"}"
  storeValues:
    token: data.token

- method: GET
  repeat: 3
  waitFor: 1000
  url: /api/geo/revgeo?latlng=12.835787029432241,80.13895099577203
  headers:
    Authorization: Bearer %token%

```

Execute the load test using the command line:

```bash
grator -config loadtest.yaml
```

## Configuration 🛠️

SeqLoad for Golang can be configured using a YAML file. Here's an overview of the configuration options in the YAML file:

- `baseURL`: The target URL or endpoint to be tested. (optional: can provide full url in actions)
- `maxActors`: The number of concurrent users or transactions to simulate.
- `actions`: An array of steps representing the sequential workflow to be tested. Each step includes:
  - `method`: HTTP method (e.g., "GET", "POST").
  - `url`: Path or endpoint of the request.
  - `body`: Optional body for the request (for methods like POST)(can be templated).
  - `waitFor`: Optional delay in milliseconds before the next step.
  - `repeat`: Optional times to repeat the step.
  - `headers`: Optional key-value pairs to be included in the request headers (can be templated).
  - `storeValues`: Optional key-value pairs to store values from the response into session values. These values can be templated into subsequent requests.

## Monitoring 📊

*Coming Soon*

## Contributing 🤝

Contributions are welcome! Whether you find a bug, want to request a feature, or contribute code, please follow the [contribution guidelines](CONTRIBUTING.md).

## License 📜

`Grators` for Golang is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

Happy load testing! 🚀
