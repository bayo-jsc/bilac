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

  <script src="node_modules/vue/dist/vue.min.js"></script>
  <script src="node_modules/axios/dist/axios.min.js"></script>
</head>
<body>
  <div id="tf">
    <div id="preloader">
      <div class="loader"></div>
    </div>

    <div class="container">
      <h2>Foosball League Table</h2>
      <div class="row">
        <div class="column">
          <h3>Tournament ${ tourID }</h3>
          <table>
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
                <td>${ team.name }</td>
                <td>${ team.PlayedMatches }</td>
                <td>${ team.GF }</td>
                <td>${ team.GA }</td>
                <td>${ team.GD }</td>
                <td>${ team.Points }</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="column">
          <h3>Matches</h3>
          <table>
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
                <td>${ findTeamWithID(match.team1ID).name }</td>
                <td>${ Math.max(0, match.team1Score) } - ${ Math.max(0, match.team2Score) }</td>
                <td>${ findTeamWithID(match.team2ID).name }</td>
                <td>
                  <button
                    class="button button-clear button-update"
                    @click="showScoreUpdate(match)"
                  >Update</button>
                </td>
              </tr>
            </tbody>
          </table>
          <button
            @click="shuffleMatch"
          >Shuffle matches</button>
        </div>
      </div>
      <a class="button" href="/draw">New Tournament</a>
    </div>

    <transition
      name="modal"
      v-if="showModal"
    >
      <div class="modal-mask">
        <div class="modal-wrapper">
          <div class="modal-container">

            <div class="modal-header">
              <slot name="header">
                Enter the score of the match
              </slot>
            </div>

            <div class="modal-body">
              <div>
                <label for="team1">${ team1Name }</label>
                <input type="number" name="score_team_1"
                  v-model="score1"
                >
              </div>

              <div>
                <label for="team1">${ team2Name }</label>
                <input type="number" name="score_team_2"
                  v-model="score2"
                  @keydown.enter="updateScore"
                >
              </div>

              <div>
                <button class="modal-default-button" @click="showModal = false">
                  Cancel
                </button>
                <button class="modal-default-button" @click="updateScore">
                  Update
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

    </transition>
  </div>


  </div>

  <small>Donate now for more future features!</small>
  <form action="https://www.paypal.com/cgi-bin/webscr" method="post" target="_top">
    <input type="hidden" name="cmd" value="_s-xclick">
    <input type="hidden" name="hosted_button_id" value="29B733CLFUC8U">
    <input type="image" src="https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif" border="0" name="submit" alt="PayPal - The safer, easier way to pay online!">
    <img alt="" border="0" src="https://www.paypalobjects.com/en_US/i/scr/pixel.gif" width="1" height="1">
  </form>

  <script src="static/js/table.min.js"></script>
</body>
</html>
