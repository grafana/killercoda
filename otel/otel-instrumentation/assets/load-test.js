import http from 'k6/http';
import { sleep } from 'k6';

const PLAYERS = [
  'Alice',
  'Bob',
  'Charlie',
  'Diana',
  'Eve',
  'Frank',
  'Grace',
  'Henry'
];

export const options = {
  vus: 5,
  duration: '5m',
};

export default function () {
  // Pick a random player
  const player = PLAYERS[Math.floor(Math.random() * PLAYERS.length)];
  
  // Make request with player name
  http.get(`http://localhost:8080/rolldice?player=${player}`);
  
  // Random sleep between 5 and 10 seconds
  sleep(Math.random() * 5 + 5);
  
  // Occasionally hit non-existent paths to generate 404s
  if (Math.random() < 0.1) {
    http.get('http://localhost:8080/nonexistent');
  }
}
