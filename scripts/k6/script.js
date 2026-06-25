// script.js
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 10,
  duration: '5s',

  // thresholds = !"зелёный"
  thresholds: {
    http_req_duration: ['p(95)<6'],  // SLO: 95% < 500ms
    http_req_failed:   ['rate<0.01'],  // < 1% ошибок
  },
};

export default function () {
  const res = http.get('http://127.0.0.1:10080/api/v1/satellite/moon');

  check(res, {
    'status 200':       (r) => r.status === 200,
    'response < 500ms': (r) => r.timings.duration < 500,
  });

  // sleep(1); //
}
