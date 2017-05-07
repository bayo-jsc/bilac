<!DOCTYPE html>
<html>
<head>
  <link rel="shortcut icon" href="favicon.ico" type="image/x-icon">

  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Bilac</title>

  <link rel="stylesheet" href="node_modules/milligram/dist/milligram.min.css">

  <script src="node_modules/vue/dist/vue.min.js"></script>
  <script src="node_modules/axios/dist/axios.min.js"></script>
</head>
<body>
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


  <small>Donate now for more future features!</small>
  <form action="https://www.paypal.com/cgi-bin/webscr" method="post" target="_top">
    <input type="hidden" name="cmd" value="_s-xclick">
    <input type="hidden" name="hosted_button_id" value="29B733CLFUC8U">
    <input type="image" src="https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif" border="0" name="submit" alt="PayPal - The safer, easier way to pay online!">
    <img alt="" border="0" src="https://www.paypalobjects.com/en_US/i/scr/pixel.gif" width="1" height="1">
  </form>

  <script src="static/js/table_football.js"></script>
</body>
</html>
