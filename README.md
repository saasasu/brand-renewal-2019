# GRIDBalance

GRIDBalance is a web application built using Go, specifically designed to demonstrate the use of Server-Sent Events (SSE). SSE is a technology enabling servers to push data to web pages over HTTP or HTTPS.

## Application Functionality

- **Data Streaming**: The application streams data every 10 seconds.
- **Timestamps**: Each data stream includes the current timestamp.
- **Forecasting**: Additionally, it provides a forecast for the following 9 intervals, updating every 10 seconds. This forecast is an example to showcase the real-time data streaming capability of the app.

## How to Test

You can test the application using a web browser or via cURL. Please note that the application may not work as expected with Firefox due to its handling of SSE.

### Using a Web Browser

1. Open a web browser (except Firefox).
2. Navigate to [https://gridbalance.fly.dev/events](https://gridbalance.fly.dev/events).
3. The application will stream data to your browser.

### Using cURL

1. Open a terminal.
2. Run the following command:

   ```sh
   curl -N https://gridbalance.fly.dev/events
   ```
   The -N flag disables buffering, which is useful for streaming data.

## About the Application

- **Technology**: Built with Go, known for its efficiency and suitability for concurrent tasks.
- **Functionality**: Utilizes Server-Sent Events (SSE) for real-time data streaming from server to client.
- **Use Case**: Demonstrates real-time streaming of energy consumption data and forecasts.

## Browser Compatibility

- The application is tested with most common web browsers. However, due to certain limitations with Firefox's handling of SSE, it is recommended to use browsers like Chrome, Safari, or Edge for the best experience.

---

Thank you for using GRIDBalance!
