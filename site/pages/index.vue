<template>
  <section class="main">
    <!-- Latest update -->
    <!-- Latest update -->
    <vue-timeline-update
      :date="new Date('2022-02-08')"
      title="12-反抗"
      description="赛季通行证皮肤：汗洛、新传奇疯玛吉、寻血猎犬威望皮肤"
      thumbnail="/images/session/session-12.jpg"
      category="announcement"
      icon="code"
      color="yellow"
    />
    <!-- Another update -->
    <vue-timeline-update
      :date="new Date('2021-11-02')"
      title="11-逃脱隐世"
      description="赛季通行证皮肤：猎手冲锋枪、新传奇艾许、新枪car、新地图风暴点、华森传家宝"
      thumbnail="/images/session/session-11.jpg"
      category="announcement"
      icon="code"
      color="turquoise"
    />

    <!-- Yet another update -->
    <vue-timeline-update
      :date="new Date('2021-08-03')"
      title="10-羽化"
      description="赛季通行证皮肤：电能冲锋枪、新传奇希尔、新枪狂暴轻机枪、竞技场排位赛、兰伯特传家宝、兰伯特改装店接管"
      thumbnail="/images/session/session-10.jpg"
      category="announcement"
      icon="code"
      color="white"
      is-last
    />
  </section>
</template>

<script>
export default {
  async asyncData({ $axios, store }) {
    store.commit('env/setCurrentNodeId', 0) // 设置当前所在node
    try {
      const [topicsPage, links] = await Promise.all([
        $axios.get('/api/topic/topics'),
        $axios.get('/api/link/toplinks'),
      ])
      console.log('topicsPage: ', topicsPage)
      console.log('links: ', links)
      return { topicsPage, links }
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

<style lang="scss" scoped></style>
