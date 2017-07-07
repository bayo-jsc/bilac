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
    <a href="/" class="column">
      <button class="button button-outline">
        Table
      </button>
    </a>

    <a href="/draw" class="column">
      <button class="button button-outline">
        Draw
      </button>
    </a>
  </div>
  <div id="tf">
    <div id="preloader">
      <div class="loader"></div>
    </div>

    <div class="container">
      <h2>Members' Elo</h2>
      
      <div class="row">
        <div class="column">
          <table>
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
              >
                <td>${ index + 1 }</td>
                <td>${ member.username }</td>
                <td>${ member.elo }</td>
                <td>${ `member.elo > 1024 ? '11' : '10'` bits }</td>
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

  <script src="static/js/elo.min.js"></script>
</body>
</html>
