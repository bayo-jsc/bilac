new Vue({
  el: "#tf",
  delimiters: ['${', '}'],

  data: {
    members: [],
    color: ['red', 'yellow', 'pink']
  },

  mounted() {
    let loader = document.getElementById("preloader")
    loader.outerHTML = ""

    this.getMembers();
  },

  methods: {
    getMembers() {
      axios.get('api/v2/members',{
        params: {
          sort: "-elo"
        }
      })
        .then(res => {
          this.members = res.data
        }, err => {
          console.log(err)
        })
    }
  }
})
