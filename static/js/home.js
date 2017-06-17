new Vue({
  el: '#tf',
  delimiters: ['${', '}'],

  data: {
    members: [],
    players: [],
  },

  mounted() {
    axios.get('api/v1/members')
      .then(res => {
        this.members = res.data
      }, err => {
        console.log(err)
      })
  },

  methods: {
    addPlayer(index) {
      this.players.push(this.members[index])
      this.members.splice(index, 1)
    },

    removePlayer(index) {
      this.members.push(this.players[index])
      this.players.splice(index, 1)
    },

    draw() {
      this.shuffleArray(this.players)
      for (let i = this.players.length - 1; i >= 0; i--) {
        console.log(this.players[i].username)
      }

    },

    createTournament() {
      return axios.post('/api/v1/tournaments', {
        teams: this.groupTeams(),
      })
        .then((res) => {
          window.location.href('/play')
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

    shuffleArray(array) {
      for (let i = array.length - 1; i >= 0; i--) {
        let j = Math.floor(Math.random() * (i + 1))
        let temp = array[i]
        array[i] = array[j]
        array[j] = temp
      }
      return array
    },
  }
})
