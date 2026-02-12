import {check} from 'k6';
import http from 'k6/http';

export const options = {
    scenarios: {
        contacts: {
            executor: 'per-vu-iterations',
            vus: 1000,
            iterations: 100,
            maxDuration: '1m',
        },
    },
};

export default function () {
    const res = http.get("http://demo.localdev.me/api/v1/users");

    check(res, {
        'Response status is 200': _ => res.status === 200
    });
}