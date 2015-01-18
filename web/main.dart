import 'package:start/start.dart';

import 'package:http_server/http_server.dart';
import 'package:logging/logging.dart';

void main() {
  Logger.root.level = Level.ALL;
  Logger.root.onRecord.listen((rec) {
    print('${rec.level.name}: ${rec.time}: ${rec.message}');
  });

  start(port: 3000).then((Server app) {

    app.static('web', jail: false);

    app.get('/lounge').listen((request) {
      request.response
        .header('Content-Type', 'text/html; charset=UTF-8')
        .send('<html><head><link rel="stylesheet" href="http://localhost:8080/stylesheets/css3clock.css"><style>body,html{margin:0;   font-family: "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif; font-weight: 300;}</style></head><body><video style="min-height:100%; min-width:100%;" autoplay><source src="http://127.0.0.1:8080/videos/uscenes_finding_nemo_fish_tank.mp4" type="video/mp4"></video><div style="position:absolute;top:1em;left:1em;opacity:0.5;text-shadow:1px 1px 4px rgba(0,0,0,0.3);color:#fff;">22E</div><div style="position:absolute;top:1em;right:1em;opacity:0.5;text-shadow:1px 1px 4px rgba(0,0,0,0.3);color:#fff;">3 loungers</div><!--clock--><div id="liveclock" class="outer_face" style="position:absolute;bottom:2em;left:44%"><div class="marker oneseven"></div><div class="marker twoeight"></div><div class="marker fourten"></div><div class="marker fiveeleven"></div><div class="inner_face"><div class="hand hour"></div><div class="hand minute"></div><div class="hand second"></div></div></div><script type="text/javascript" src="http://localhost:8080/javascripts/css3clock.js"></script><!--/clock--></body></html>');
    });

    app.get('/pod').listen((request) {
      request.response
        .header('Content-Type', 'text/html; charset=UTF-8')
        .send('<html><head><style>body,html{margin:0;   font-family: "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif; font-weight: 300;}</style></head><body>You have the power.  Pair.</body></html>');
    });
    
    app.get('/hello/:name.:lastname?').listen((request) {
      request.response
        .header('Content-Type', 'text/html; charset=UTF-8')
        .send('Hello, ${request.param('name')} ${request.param('lastname')}');
    });

    app.ws('/socket').listen((socket) {
      socket.on('connected').listen((data) {
        socket.send('ping', 'data-from-ping');
      });

      socket.on('pong').listen((data) {
        print('pong: $data');
        socket.close(1000, 'requested');
      });

      socket.onOpen.listen((ws) {
        print('new socket opened');
      });

      socket.onClose.listen((ws) {
        print('socket has been closed');
      });
    });
  });
}