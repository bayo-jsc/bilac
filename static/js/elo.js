new Vue({
  el: "#tf",
  delimiters: ['${', '}'],

  data: {
    members: [],
    color: ['#4dd0e1', '#80deea', '#b2ebf2', '#e0f7fa']
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
