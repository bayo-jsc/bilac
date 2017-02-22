new Vue({
  el: '#app',
  delimiters: ['${', '}'],

  data() {
    return {
      membersName: [
        'dungpt',
        'hiepph',
        'anhmt',
        'bangcht',
        'anhnq',
        'khainv',
        'dungnt',
        'hoannx',
      ],
      members: [],
      newMem: {},
      newMember: '',
    }
  },

  computed: {
    membersList() {
      return this.membersName.map((member, index) => ({
        id: index,
        username: member,
        team_id: 0,
      }))
    }
  },

  watch: {
    newMember(value) {
      if (value !== '') {
        axios.post('api/v1/members', {

        })
      }
    }
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
    createMember() {
      axios.post('/api/v1/members', {
        username: this.newMem.username
      }).then(res => {
          this.members.push(res.data)
          this.newMem = {}
        }, err => {
          console.log(err)
        })
    },

    destroyMember(index) {
      axios.delete('/api/v1/members/' + this.members[index].id)
        .then(res => {
          this.members.splice(index, 1)
        }, err => {
          console.log(err)
        })
    },

    getMembers() {
      axios
      .get('/api/v1/members')
      .then(response => {
        this.$set(this, 'membersList', response.data)
      })
    },

    draw() {
      axios.patch('/api/v1/draw')
        .then(res => {
          this.members = res.data
        }, err => {
          console.log(err)
        })
    },
  }
})
