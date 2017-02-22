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
    <select multiple v-model="members">
      <option v-for="member in membersList" :value="member">${ member }</option>
    </select>
  
    <table id='tf'>
      <thead>
        <tr>
          <th>Team</th>
          <th>Username</th>
          <th></th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="(mem, index) in members">
          <td>${ mem.team_id }</td>
          <td>${ mem.username }</td>
          <td>
            <button class="button button-clear" v-on:click="destroyMember(index)">x Remove</button>
          </td>
        </tr>
        <tr>
          <td>#</td>
          <td>
            <input type="text"
              placeholder="username"
              v-model="newMem.username"
              v-on:keyup.enter="createMember">
          </td>
          <td>
            <button class="button button-default" type="button" v-on:click="createMember">+ Add</button>
          </td>
        </tr>
        <tr>
          <td>
            <button class="button button-default" type="button" v-on:click="draw">DRAW</button>
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
