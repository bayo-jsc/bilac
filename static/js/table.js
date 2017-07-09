new Vue({
  el: "#tf",
  delimiters: ['${', '}'],

  data: {
    tourIDs: [],
    tourID: 0,
    lastTourID: 0,
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

    this.getTourIDs()
  },

  watch: {
    tourID() {
      this.getTournament()
    },
  },

  methods: {
    getTourIDs() {
      axios.get('/api/v2/tournaments')
        .then(res => {
          let ids = res.data.map(tour => tour.ID)

          this.$set(this, 'tourIDs', ids)
          this.$set(this, 'tourID', ids[0])
          this.$set(this, 'lastTourID', ids[0])
          this.getTournament()
        }, err => {
          console.log(err)
        })
    },

    getTournament() {
      axios.get(`/api/v2/tournaments/${this.tourID}`)
        .then(res => {
          const data = res.data

          this.$set(this, 'teams', data.Teams)
          this.$set(this, 'matches', data.Matches)

        }, err => {
          console.log(err)
        })
    },

    matchScore(team1, team2) {
      let game = this.matchAt(team1, team2)

      return game.team1ID ? game.team1Score + "-" + game.team2Score : "--"
    },

    findTeamWithID(teamID) {
      return this.teams.find(x => x.ID === teamID)
    },

    showScoreUpdate(match) {
      this.$set(this, 'team1Name', this.teamName(this.findTeamWithID(match.Team1ID)))
      this.$set(this, 'team2Name', this.teamName(this.findTeamWithID(match.Team2ID)))

      this.$set(this, 'score1', Math.max(0, match.Team1Score))
      this.$set(this, 'score2', Math.max(0, match.Team2Score))
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

    teamName(team) {
      return `${ team.Member1.username } + ${ team.Member2.username }`
    },

    team1NameWithElo(match) {
      const team = this.findTeamWithID(match.Team1ID)
      const mem1Change = match.Mem1EloAfter - match.Mem1EloBefore
      const mem2Change = match.Mem2EloAfter - match.Mem2EloBefore
      return `${ team.Member1.username }(${mem1Change}) 
             ${ team.Member2.username }(${mem2Change})`
    },

    team2NameWithElo(match) {
      const team = this.findTeamWithID(match.Team2ID)
      const mem3Change = match.Mem3EloAfter - match.Mem3EloBefore
      const mem4Change = match.Mem4EloAfter - match.Mem4EloBefore
      return `${ team.Member1.username }(${mem3Change}) 
             ${ team.Member2.username }(${mem4Change})`
    },
  },
})
