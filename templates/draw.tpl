<!DOCTYPE html>
<html>
<head>
  <link rel="shortcut icon" href="favicon.ico" type="image/x-icon">

  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Bilac</title>

  <link rel="stylesheet" href="node_modules/milligram/dist/milligram.min.css">
  <link rel="stylesheet" href="static/css/app.css">

  <script src="node_modules/vue/dist/vue.js"></script>
  <script src="node_modules/axios/dist/axios.min.js"></script>
</head>
<body>
  <div class="row">
    <div class="column">
      Go to:
    </div>
    <a href="/">
      <button class="column button button-outline">
        Table
      </button>
    </a>

    <a href="/elo">
      <button class="column button button-outline">
        Elo
      </button>
    </a>
  </div>

  <div id="tf">
    <div id="preloader">
      <div class="loader"></div>
    </div>

    <div class="container">
      <div class="row">
        <div class="column">
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>Username</th>
                <th></th>
              </tr>
            </thead>

            <tbody>
              <tr
                v-for="mem, index in members"
              >
                <td>${ mem.ID }</td>
                <td>${ mem.username }</td>
                <td>
                  <button class="button button-clear" v-on:click="addPlayer(index)">+ Add</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="column">
          <table>
            <thead>
              <tr>
                <th>Team ID</th>
                <th>Username</th>
                <th></th>
              </tr>
            </thead>

            <tbody>
              <tr
                v-for="player, index in players"
                :key="player.id"
              >
                <td>${ Math.trunc(index / 2) + 1 }</td>
                <td>${ player.username }</td>
                <td>
                  <button class="button button-clear" v-on:click="removePlayer(index)">x Remove</button>
                </td>
              </tr>
              <tr>
                <td>
                  <button class="button button-default" type="button" v-on:click="draw">Draw</button>
                  <button class="button button-default" type="button" v-on:click="createTournament">Create Tournament</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>

  <small>Donate now for more future features!</small>
  <form action="https://www.paypal.com/cgi-bin/webscr" method="post" target="_top">
    <input type="hidden" name="cmd" value="_s-xclick">
    <input type="hidden" name="hosted_button_id" value="29B733CLFUC8U">
    <input type="image" src="https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif" border="0" name="submit" alt="PayPal - The safer, easier way to pay online!">
    <img alt="" border="0" src="https://www.paypalobjects.com/en_US/i/scr/pixel.gif" width="1" height="1">
  </form>

  <script src="static/js/draw.min.js"></script>
</body>
</html>
