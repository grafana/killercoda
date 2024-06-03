# Instrumenting Your React Application

In this section, we'll instrument the React application to begin sending data to your Grafana Cloud instance. This process is incredibly quick and easy and provides out of the box instrumentation for all of the Core Web Vitals, as well as user sessions, errors, and events.

## Step 1: Installing the Grafana Faro SDK

Start by navigating to the `frontend` folder if you're not already there.

```bash
cd frontend
```

Next, install the `faro-react` and `faro-web-tracing` packages from NPM.

```bash
npm i -S @grafana/faro-react @grafana/faro-web-tracing
```{{execute}}

## Step 2: Adding the Instrumentation Code

Now that the packages have been installed, it's time to instrument our application by adding the necessary initialization code.

Open the `index.js` file in a text editor such as `nano`.

```bash
nano src/index.js
```{{execute}}

At the top of the file, add the following imports:

```javascript
import { createRoutesFromChildren, matchRoutes, Routes, useLocation, useNavigationType } from 'react-router-dom';
import { getWebInstrumentations, initializeFaro, ReactIntegration, ReactRouterVersion } from '@grafana/faro-react';
import { TracingInstrumentation } from '@grafana/faro-web-tracing';
```{{copy}}

Next, directly after the import statements and *before* the React root element is created, add the Faro initialization code. Be sure to enter **your** `<COLLECTOR_URL>` from the **Web SDK Configuration** section in your Grafana Frontend Observability app.

```javascript
var faro = initializeFaro({
  url: '<COLLECTOR_URL>',
  app: {
    name: 'fo11y-demo',
    version: '1.0.0',
    environment: 'production'
  },
  instrumentations: [
    ...getWebInstrumentations(),
    new TracingInstrumentation(),
    new ReactIntegration({
        router: {
            version: ReactRouterVersion.V6,
            dependencies: {
                createRoutesFromChildren,
                matchRoutes,
                Routes,
                useLocation,
                useNavigationType,
            },
        }
    })
  ],
});
```{{copy}}

Now save the file and exit `nano` (*You can do this by pressing CTRL+X on your keyboard, hitting Y and then Enter to save the changes*).

Next, open the `src/App.js` file in `nano`.

```bash
nano src/App.js
```{{execute}}

With the `App.js` file open, import the `FaroRoutes` component from `@grafana/faro-react`.

```javascript
import { FaroRoutes } from '@grafana/faro-react';
```{{copy}}

Next, remove *just* the `Routes` import from the `react-router-dom` import statement as we won't be using this component now.

Now, replace the existing `Routes` component with the `FaroRoutes` component. Your `<Routes>` section should now look like the following:

```javascript
<FaroRoutes>
    <Route path="/" element={<Home />} />
    <Route path="/favorites" element={<Favorites />} />
    <Route path="/search" element={<Search />} />
    <Route path="/about" element={<About />} />
</FaroRoutes>
```{{copy}}

Your React application is now instrumented with Grafana Faro.

Re-run the application by navigating back to the root `fo11y-demo` directory and re-running the containers with docker compose. *Remember to take the containers down first with `docker-compose down`*.

```bash
cd ..
docker-compose down
docker-compose up --build -d
```{{execute}}

Now, open the frontend at [http://locahost:3000]({{TRAFFIC_HOST1_3000}}) and click around the application to generate some data.

Now visit your Grafana Cloud Frontend Observability app and you should begin to see data from your application!