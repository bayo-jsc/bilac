new Vue({
  el: '#tf',
  delimiters: ['${', '}'],

  data: {
    members: [],
    players: [],
  },

  mounted() {
    let loader = document.getElementById("preloader")
    loader.outerHTML = ""

    this.getMembers()
  },

  methods: {
    getMembers() {
      axios.get('api/v2/members')
        .then(res => {
          this.members = res.data
        }, err => {
          console.log(err)
        })
    },

    addPlayer(index) {
      this.players.push(this.members[index])
      this.members.splice(index, 1)
    },

    removePlayer(index) {
      this.members.push(this.players[index])
      this.players.splice(index, 1)
    },

    draw() {
      this.randomList(this.players)
    },

    createTournament() {
      return axios.post('/api/v2/tournaments', {
        teams: this.groupTeams(),
      })
        .then((res) => {
          window.location.href = '/'
        }, (err) => {
          console.log(err)
        })
    },

    groupTeams() {
      let teams = new Array(Math.trunc(this.players.length / 2))
      for (var i = 0; i < teams.length; i++) {
        teams[i] = {
          member1_id: this.players[i * 2].ID,
          member2_id: this.players[i * 2 + 1].ID
        }
      }
      return teams
    },

    randomList(rand) {
      return rand.sort(() => { return 0.5 - Math.random() });
    },
  }
})