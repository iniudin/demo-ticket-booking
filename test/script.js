import http from 'k6/http';
import {check, sleep} from 'k6';

export const options = {
    stages: [
        {duration: '30s', target: 20},
        {duration: '1m30s', target: 10},
        {duration: '20s', target: 0},
    ],
    thresholds: {
        http_req_failed: ['rate<0.02'],
        http_req_duration: ['p(95)<2000']
    },
    ext: {
        loadimpact: {
            // Project: Default project
            projectID: 3650153,
            // Test runs with the same name groups test runs together
            name: 'YOUR TEST NAME'
        }
    }

};

export default function () {
    const res = http.get('http://localhost:8080/api/v1/concerts');
    check(res, {'status was 200': (r) => r.status == 200});
    sleep(1);
}
