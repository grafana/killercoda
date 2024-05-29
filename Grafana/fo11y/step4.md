# Enhancing Application Observability with Custom Metrics, Events, and Logs

In this step, we'll dive deeper into how to enrich the observability of your React application by sending custom metrics, events, logs, and errors using the Grafana Faro SDK. This level of detailed monitoring allows you to capture the nuances of user interactions and system performance, providing valuable insights into your application's operation and user experiences.

## Sending Custom Log Entries

By default, the Faro SDK hooks into the browser's console object to send info, warning, and error logs to your Grafana Cloud. However, you can also manually send your own log entries directly to the collector. This capability is particularly useful for troubleshooting and gaining additional context about user activities within your application.

### Log User Search Activities

When a user performs a search in your application, you can log this action along with the results.

To do this, open up `src/components/Search.js` using `nano`:

```bash
nano frontend/src/components/Search.js
```

Import the `faro` object:

```javascript
import { faro } from '@grafana/faro-react';
```

Then add the following just after the `setGames(response.data)` line:

```javascript
faro.api.pushLog([`Search result for ${searchTerm} found ${response.data.length} games.`], {
	level: 'warn',
	context: {
		searchTerm: searchTerm,
		results: response.data.length,
		userId: userUUID
	}
});
```

This custom log entry will help you monitor how users interact with the search feature and track the effectiveness of your search functionality.

## Sending Custom Events

Events are similar to logs but are specifically intended as indicators of user interactions within your application. They provide a structured way to track how users engage with different features.

### Track Search and Favorites Interactions

Let's also push an event to Grafana whenever the user performs a search. Add the following right after our `pushLog` call.

```javascript
faro.api.pushEvent('search', { userId: userUUID, searchTerm: searchTerm }, 'search');
```

We can also track when users favourite/unfavourite games.

Open up the `src/components/GameCard.js` file in `nano`:

```bash
nano frontend/src/components/GameCard.js
```

Import the `faro` object:

```javascript
import { faro } from '@grafana/faro-react';
```

After `setIsFavorited(false);`:

```javascript
faro.api.pushEvent('unfavorited', { user_id: userUUID, game_id: `${game.id}` }, 'favorites');
```

After `setIsFavorited(true);`:

```javascript
faro.api.pushEvent('favorited', { user_id: userUUID, game_id: `${game.id}` }, 'favorites');
```

## Sending Custom Errors

While Faro automatically captures unhandled errors, there are scenarios where you might want to manually report errors, especially those caught in try-catch blocks within your application.

### Manually Report a Search Failure

Open up `src/components/Search.js` again in `nano` and add the line inside the `catch` block:

```bash
nano frontend/src/components/Search.js
```

```javascript
faro.api.pushError(error);
```

## Restart the application

To test out these changes, restart the application via Docker:

```bash
docker-compose down
docker-compose up --build -d
```

Open your application again by visiting [localhost:3000]({{TRAFFIC_HOST1_3000}}) and generate some logs and events by performing searches as well as favoriting/unfavoriting games.

You should now be able to observe the logs and events within the Grafana Cloud Frontend Observability application. 

For a detailed demonstration of this in action, please refer to the [YouTube video found here](https://www.youtube.com/watch?v=IA_-zkpVhIU).

## Conclusion

By implementing these custom metrics, events, logs, and error capturing, you significantly enhance your ability to observe and understand both the technical performance and user experience of your React application. Continue exploring these tools to optimize your monitoring setup and improve your application.