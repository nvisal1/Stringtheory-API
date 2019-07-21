const singleton = Symbol();
const singletonEnforcer = Symbol();
const url = 'localhost:8080';

class SocketDriver {
    constructor(enforcer) {
        if (enforcer !== singletonEnforcer) {
            throw new Error('Cannot construct singleton');
        }
        this.connect(url)
    }

    connect(url) {
        this.websocket = new WebSocket('ws://' + url + "/record")
    }

    static get instance() {
        if (!this[singleton]) {
            this[singleton] = new SocketDriver(singletonEnforcer);
        }
        return this[singleton];
    }
}
