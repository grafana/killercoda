import http from 'k6/http';
import { sleep } from 'k6';
import { randomItem } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

const players = ['John', 'Jane', 'Bert', 'Ed'];

export const options = {
  // A number specifying the number of VUs to run concurrently.
  vus: 2,
  // A string specifying the total duration of the test run.
  duration: '60m',
};

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function() {
  const randomPlayer = randomItem(players);
  http.get(`http://localhost:8080/rolldice?player=${randomPlayer}`);
  sleep(Math.random() * 10);
}
