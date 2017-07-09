<!DOCTYPE html>
<html>
<head>
  <link rel="shortcut icon" href="favicon.ico" type="image/x-icon">

  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Bilac</title>

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
        <div class="col s10">
          <h2>Bayo Bilac League</h2>
        </div>
        <div class="col s2">
          <label>Select tournament</label>
          <select
            v-model="tourID"
          >
            <option
                v-for="id in tourIDs"
                :value="id"
            >
                ${ id }
            </option>
          </select>
        </div>
      </div>

      <div class="row">
        <div class="col m6">
          <h5>Tournament ${ tourID }</h5>
          <table class="bordered">
            <thead>
              <tr>
                <th>Rank</th>
                <th>Team</th>
                <th>Played</th>
                <th>GF</th>
                <th>GA</th>
                <th>GD</th>
                <th>Points</th>
              </tr>
            </thead>

            <tbody>
              <tr
                v-for="team, index in teams"
              >
                <td>${ index + 1 }</td>
                <td>${ teamName(team) }</td>
                <td>${ team.PlayedMatches }</td>
                <td>${ team.GF }</td>
                <td>${ team.GA }</td>
                <td>${ team.GD }</td>
                <td>${ team.Points }</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="col m6">
          <h5>Matches</h5>
          <table class="striped">
            <thead>
              <tr>
                <th></th>
                <th>Score</th>
                <th></th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="match, index in matches"
              >
                <td>
                  ${ team1NameWithElo(match) }
                </td>
                <td>${ evalScore(match.Team1Score) } - ${ evalScore(match.Team2Score) }</td>
                <td>${ team2NameWithElo(match) }</td>
                <td>
                  <button
                    class="waves-effect waves-light btn"
                    @click="showScoreUpdate(match)"
                    v-if="tourID === lastTourID"
                  >Update</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <a class="waves-effect waves-light btn" href="/draw">New Tournament</a>

      <br>
      <small>Donate now for more future features!</small>
      <form action="https://www.paypal.com/cgi-bin/webscr" method="post" target="_top">
        <input type="hidden" name="cmd" value="_s-xclick">
        <input type="hidden" name="hosted_button_id" value="29B733CLFUC8U">
        <input type="image" src="https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif" border="0" name="submit" alt="PayPal - The safer, easier way to pay online!">
        <img alt="" border="0" src="https://www.paypalobjects.com/en_US/i/scr/pixel.gif" width="1" height="1">
      </form>
    </div>

    <transition
      name="modal"
      v-if="showModal"
    >
      <div class="modal-mask">
        <div class="modal-wrapper">
          <div class="modal-container">

            <div class="modal-header">
              <h4>
                Enter the score of the match
              </h4>
            </div>

            <div class="modal-body">
              <div>
                <label for="team1">${ team1Name }</label>
                <input
                  v-model="score1"
                  type="number"
                  name="score_team_1"
                  min="0"
                  max="10"
                >
              </div>

              <div>
                <label for="team1">${ team2Name }</label>
                <input
                  v-model="score2"
                  @keydown.enter="updateScore"
                  type="number"
                  name="score_team_2"
                  min="0"
                  max="10"
                >
              </div>

              <div class="row">
                <div class="col s6">
                  <button class="waves-effect waves-light brown darken-1 btn" @click="showModal = false">
                    Cancel
                  </button>
                </div>
                <div class="col s6">
                  <button class="waves-effect waves-light btn" @click="updateScore">
                    Update
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

    </transition>
  </div>


  </div>

  <script src="static/js/table.min.js"></script>
</body>
</html>
