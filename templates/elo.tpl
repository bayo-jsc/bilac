<!DOCTYPE html>
<html>
<head>
  <link rel="shortcut icon" href="static/favicon.ico" type="image/x-icon">

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
    {{ template "navbar" .}}

    <div id="preloader">
      <div class="loader"></div>
    </div>

    <div class="container">
      <h3>Members' Elo</h3>
      
      <div class="row">
        <div class="col s12">
          <table class="striped">
            <thead>
              <tr>
                <th>Rank</th>
                <th>Member</th>
                <th>Elo</th>
                <th>Bit</th>
              </tr>
            </thead>

            <tbody>
              <tr
                v-for="member, index in members"
                :style="{ 'background-color': color[index] }"
              >
                <td v-if="index !== 0">${ index + 1 }</td>
                <td v-else>
                  <img src="https://elearningimages.adobe.com/files/2011/05/First.jpg" height="50px">
                </td>
                <td>${ member.username }</td>
                <td>${ member.elo }</td>
                <td>${ member.elo > 1023 ? 11 : 10 }</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      {{ template "donate" }}
    </div>
  </div>


  <script src="static/js/elo.min.js"></script>
</body>
</html>
