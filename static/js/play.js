Vue.component('modal', {
  template: '#modal-template'
})

new Vue({
  el: "#tf",
  delimiters: ['${', '}'],

  data: {
    tourID: 0,
    matches: [],
    teams: [],
    showModal: false,
  },

  mounted() {
    this.getTournament()
  },

  methods: {
    getTournament() {
      axios.get('/api/v1/last-tournament')
        .then(res => {
          const data = res.data

          let matches = data.Matches.map(match => {
            return {
              matchID: match.ID,
              team1ID: match.Team1ID,
              team2ID: match.Team2ID,
              team1Score: match.Team1Score,
              team2Score: match.Team2Score,
            }
          })

          this.$set(this, 'tourID', data.ID)
          this.$set(this, 'matches', matches)

          this.getTeams()
        }, err => {
          console.log(err)
        })
    },

    getTeams() {
      axios.get('api/v1/tournaments/' + this.tourID + '/teams')
        .then(res => {
          const data = res.data

          let teams = data.map(team => {
            return {
              ID: team.ID,
              name: team.Member1.username + "+" + team.Member2.username,
            }
          })

          this.$set(this, 'teams', teams)
        }, err => {
          console.log(err)
        })
    },

    matchScore(team1, team2) {
      let game = this.matchAt(team1, team2)

      // console.log(game)
      return game.team1ID ? game.team1Score + "-" + game.team2Score : "--"
    },

    matchAt(team1, team2) {
      for (let match of this.matches) {
        if (match.team1ID == team1.ID && match.team2ID == team2.ID) {
          return match
        }
      }
      return {}
    },

    updateScore(team1, team2) {
      let match = this.matchAt(team1, team2)


    }
  },
})
