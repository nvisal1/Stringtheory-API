

const config = {
	type: Phaser.AUTO,
	// width: window.screen.availWidth * window.devicePixelRatio,
	// height: window.screen.availHeight * window.devicePixelRatio,
	width: 1000,
	height: 500,
	scene: {
		preload: preload,
		create: create,
		update: update,
	}
};

const notes = ['A', 'B', 'C', 'D', 'E', 'F', 'G'];

let websocket;
let text;
let score = 0;
let currentNote;
const game = new Phaser.Game(config);

function preload() {

}

function create() {
	websocket = SocketDriver.instance.websocket;
	const style = { font: "bold 32px Arial", fill: "#fff", boundsAlignH: "center", boundsAlignV: "middle" };
	currentNote = notes[Math.floor(Math.random() * (notes.length - 0))]
	noteText = this.add.text(0, 0, currentNote, style);
	scoreText = this.add.text(80, 0, `Score: ${score}`, style);
	timedEvent = this.time.addEvent({delay: 2000, callback: selectNote, callbackScope: this, repeat: 10});
}

function update() {
	websocket.onmessage = function(e) {
		if (JSON.parse(e.data).Note === currentNote) {
			score++;
		}
	};
}

function selectNote() {
	currentNote = notes[Math.floor(Math.random() * (notes.length - 0))]
	noteText.setText(currentNote);
	scoreText.setText(score);
}
