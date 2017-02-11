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
  <table id='tf'>
    <thead>
      <tr>
        <th>ID</th>
        <th>Username</th>
      </tr>
    </thead>

    <tbody>
      <tr v-for="mem in members">
        <td>${ mem.id }</td>
        <td>${ mem.username }</td>
      </tr>
    </tbody>
  </table>

  <script src="static/js/table_football.js"></script>
</body>
</html>
