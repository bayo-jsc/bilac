<!DOCTYPE html>
<html>
<head>
  <link rel="shortcut icon" href="static/favicon.ico" type="image/x-icon">

  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Bilac</title>

  <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
  <link rel="stylesheet" href="node_modules/materialize-css/dist/css/materialize.min.css">
  <link rel="stylesheet" href="static/css/app.css">

  <script src="node_modules/vue/dist/vue.js"></script>
  <script src="node_modules/axios/dist/axios.min.js"></script>
</head>
<body>
  <div id="tf">
    <div id="preloader">
      <div class="loader"></div>
    </div>

    {{ template "navbar" . }}

    <div class="container">
      <div class="row">
        <h3>Choose player to draw</h3>
      </div>
      <div class="row">
        <div class="col m6">
          <table class="striped">
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
                  <button class="waves-effect waves-light green btn" v-on:click="addPlayer(index)">
                    Add<i class="material-icons right">add</i>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="col m6">
          <table class="striped">
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
                  <button class="waves-effect waves-light red btn" v-on:click="removePlayer(index)">
                    Remove<i class="material-icons right">delete</i>
                  </button>
                </td>
              </tr>
              <tr>
                <td>
                  <button
                    v-if="players.length > 0"
                    v-on:click="draw"
                    class="waves-effect waves-light btn"
                    type="button"
                  >
                    Draw
                  </button>
                </td>
                <td>
                  <button
                    v-if="players.length > 0 && isDrawed"
                    class="waves-effect waves-light btn"
                    type="button"
                    v-on:click="createTournament"
                  >
                    Create
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <br>
      {{ template "donate" }}
    </div>
  </div>

  <script src="static/js/draw.min.js"></script>
</body>
</html>
