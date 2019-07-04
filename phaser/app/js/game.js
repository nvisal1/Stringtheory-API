

const config = {
	type: Phaser.AUTO,
	width: window.screen.availWidth * window.devicePixelRatio,
	height: window.screen.availHeight * window.devicePixelRatio,
	scene: {
		preload: preload,
		create: create,
		update: update,
	}
};

let websocket;
let text;
const game = new Phaser.Game(config);

function preload() {

}

function create() {
	websocket = SocketDriver.instance.websocket;
	websocket.onopen = () => websocket.send({ Note: 'V'});
	websocket.onopen = () => websocket.send(JSON.stringify({ Note: 'C'}));
	websocket.onopen = () => websocket.send(JSON.stringify({ Note: 'D'}));
	text = this.add.text(0, 0, 'A');
}

function update() {
	websocket.onmessage = function(e){ text.setText(JSON.parse(e.data).Note); };
}
