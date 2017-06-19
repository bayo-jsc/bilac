new Vue({
  el: "#tf",
  delimiters: ['${', '}'],

  data: {
    tourID: 0,
    matches: [],
    teams: [],
    team1Name: "",
    team2Name: "",
    matchID: 0,
    score1: 0,
    score2: 0,
    showModal: false,
  },

  mounted() {
    let loader = document.getElementById("preloader")
    loader.outerHTML = ""

    this.getTournament()
  },

  methods: {
    getTournament() {
      axios.get('/api/v2/last-tournament')
        .then(res => {
          const data = res.data

          let matches = data.Matches.map(match => {
            return {
              ID: match.ID,
              team1ID: match.Team1ID,
              team2ID: match.Team2ID,
              team1Score: match.Team1Score,
              team2Score: match.Team2Score,
            }
          })

          let teams = data.Teams.map(team => {
            return {
              ID: team.ID,
              name: team.Member1.username + " + " + team.Member2.username,
              GF: team.GF,
              GA: team.GA,
              GD: team.GD,
              Points: team.Points,
              PlayedMatches: team.PlayedMatches,
            }
          })
          this.$set(this, 'tourID', data.ID)
          this.$set(this, 'teams', teams)
          this.$set(this, 'matches', matches)

        }, err => {
          console.log(err)
        })
    },

    matchScore(team1, team2) {
      let game = this.matchAt(team1, team2)

      return game.team1ID ? game.team1Score + "-" + game.team2Score : "--"
    },

    findTeamWithID(teamID) {
      return this.teams.find(x => x.ID == teamID)
    },

    showScoreUpdate(match) {
      this.$set(this, 'team1Name', this.findTeamWithID(match.team1ID).name)
      this.$set(this, 'team2Name', this.findTeamWithID(match.team2ID).name)

      this.$set(this, 'score1', match.team1Score)
      this.$set(this, 'score2', match.team2Score)
      this.$set(this, 'matchID', match.ID)
      this.$set(this, 'showModal', true)
    },

    updateScore() {
      axios.patch('/api/v2/tournaments/' + this.tourID + '/matches/' + this.matchID, {
        score_team_1: this.score1,
        score_team_2: this.score2,
      })
        .then(res => {
          this.getTournament()
          this.$set(this, 'showModal', false)
        }, err => {
          console.log(err)
        })
    },

    shuffleMatch() {
      axios.patch('/api/v2/tournaments/' + this.tourID + '/shuffle')
        .then(res => {
          this.getTournament()
        }, err => {
          console.log(err)
        })
    }
  },
})
