import http from 'k6/http';
import { check } from 'k6';

export const options = {
  scenarios: {
    users: {
      executor: 'per-vu-iterations',
      vus: 2000,
      iterations: 10000,
      maxDuration: '30s',
    },
  },
};

export default function () {
  const res = http.get('http://demo.localdev.me/api/v1/users');
  check(res, { 'status was 200': (r) => r.status === 200 });
}
