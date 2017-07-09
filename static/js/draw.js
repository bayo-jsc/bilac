new Vue({
  el: '#tf',
  delimiters: ['${', '}'],

  data: {
    members: [],
    players: [],
    isDrawed: false,

  },

  mounted() {
    let loader = document.getElementById("preloader")
    loader.outerHTML = ""

    this.getMembers()
  },

  methods: {
    getMembers() {
      axios.get('api/v2/members', {
        params: {
          sort: "ID",
        }
      })
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
      this.players = this.shuffle(this.players)
      this.$set(this, 'isDrawed', true)
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

    randomList(array) {
      return array.sort(() => { return 0.5 - Math.random() });
    },

    shuffle(arr) {
      var array = arr.slice()
      var currentIndex = array.length,
          temporaryValue,
          randomIndex

      // While there remain elements to shuffle...
      while (0 !== currentIndex) {

        // Pick a remaining element...
        randomIndex = Math.floor(Math.random() * currentIndex)
        currentIndex -= 1

        // And swap it with the current element.
        temporaryValue = array[currentIndex]
        array[currentIndex] = array[randomIndex]
        array[randomIndex] = temporaryValue
      }

      return array
    },
  }
})
