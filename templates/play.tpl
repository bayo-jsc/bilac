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
    <div class="container">
      <div class="row">
        <div class="column">
          <table>
            <thead>
              <tr>
                <th>Team</th>
                <template
                  v-for="team, index in teams"
                >
                  <th>${ team.name }</th>
                </template>
              </tr>
            </thead>

            <tbody>
              <tr
                v-for="team, index1 in teams"
              >
                <td>${ team.name }</td>
                <template
                  v-for="otherTeam, index2 in teams"
                >
                  <td
                    :click="updateScore(team, otherTeam)"
                  >
                    ${ index1 < index2 ? matchScore(team, otherTeam) : "--" }
                  </td>
                </template>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

     <button id="show-modal" @click="showModal = true">Show Modal</button>
  <!-- use the modal component, pass in the prop -->
  <modal v-if="showModal" @close="showModal = false">
    <!--
      you can use custom content here to overwrite
      default content
    -->
    <h3 slot="header">custom header</h3>
  </modal>
  </div>

  <small>Donate now for more future features!</small>
  <form action="https://www.paypal.com/cgi-bin/webscr" method="post" target="_top">
    <input type="hidden" name="cmd" value="_s-xclick">
    <input type="hidden" name="hosted_button_id" value="29B733CLFUC8U">
    <input type="image" src="https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif" border="0" name="submit" alt="PayPal - The safer, easier way to pay online!">
    <img alt="" border="0" src="https://www.paypalobjects.com/en_US/i/scr/pixel.gif" width="1" height="1">
  </form>

  
  <script type="text/x-template" id="modal-template">
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">

          <div class="modal-header">
            <slot name="header">
              default header
            </slot>
          </div>

          <div class="modal-body">
            <slot name="body">
              default body
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              default footer
              <button class="modal-default-button" @click="$emit('close')">
                OK
              </button>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
  </script>

  <script src="static/js/play.js"></script>
</body>
</html>
