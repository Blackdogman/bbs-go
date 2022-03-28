<template>
  <section class="main">
    <h1>HelloWorld</h1>
  </section>
</template>

<script>
export default {
  async asyncData({ $axios, store }) {
    store.commit('env/setCurrentNodeId', 0) // 设置当前所在node
    try {
      const [nodes, topicsPage, scoreRank, links] = await Promise.all([
        $axios.get('/api/topic/nodes'),
        $axios.get('/api/topic/topics'),
        $axios.get('/api/user/score/rank'),
        $axios.get('/api/link/toplinks'),
      ])
      return { nodes, topicsPage, scoreRank, links }
    } catch (e) {
      console.error(e)
    }
  },
  data() {},
  head() {
    return {
      title: this.$siteTitle(),
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: this.$siteDescription(),
        },
        { hid: 'keywords', name: 'keywords', content: this.$siteKeywords() },
      ],
    }
  },
}
</script>

<style lang="scss" scoped>
  h1 {
    font-size: 72px;
  }
</style>
