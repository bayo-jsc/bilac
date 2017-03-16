<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Duel</title>

  <link rel="stylesheet" href="node_modules/milligram/dist/milligram.min.css">

  <script src="node_modules/vue/dist/vue.min.js"></script>
  <script src="node_modules/axios/dist/axios.min.js"></script>
</head>
<body>
  <div id="app">
    Select players:
    <select multiple v-model="players" style="height: auto">
      <option v-for="member in members" :value="member">${ member.username }</option>
    </select>

    <div>
      Create new member:
      <input type="text" v-model="newMember" @keydown.enter="createMember">
      <button @click="createMember">Create member</button>
    </div>

    <table id='tf' v-show="!isShowScoreboard">
      <thead>
        <tr>
          <th>Team</th>
          <th>Username</th>
          <th></th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="(mem, index) in players">
          <td>${ mem.team_id }</td>
          <td>${ mem.username }</td>
          <td>
            <button class="button button-clear" v-on:click="destroyMember(index)">x Remove</button>
          </td>
        </tr>
        <tr>
          <td>
            <button class="button button-default" type="button" v-on:click="draw">DRAW</button>
          </td>
        </tr>
      </tbody>
    </table>

    <table id="scoreboard" v-show="isViewScoreboard">
      <thead>
        <tr>
          <th>Team</th>
          <th>Point</th>
          <th>Score</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="(team, index) in teams">
          <td>
            ${ team.id }
            <span v-for="mem in team.members">
              ${ mem.username }
            </span>
          </td>
          <td>
            ${ mem.point }
          </td>
          <td>
            ${ mem.score }
          </td>
        </tr>
      </tbody>
    </table>
  </div>

  <div>
    Rule:
    <ol>
      <li>Each match is 5 minutes</li>
      <li>Last position team lost 10k VND/each member</li>
      <li>Second position from bottom lost 5k VND/each member</li>
    </ol>
  </div>

  <script src="static/js/table_football.js"></script>
</body>
</html>
