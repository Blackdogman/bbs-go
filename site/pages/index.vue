<template>
  <section class="main">
    <div v-for="(session, index) in apexSessions" :key="index">
      <vue-timeline-update
        v-if="index !== apexSessions.length - 1"
        :date="session.date"
        :title="session.title"
        :description="session.description"
        :thumbnail="session.thumbnail"
        category="announcement"
        icon="code"
        :color="session.color"
      />
      <vue-timeline-update
        v-if="index === apexSessions.length - 1"
        :date="session.date"
        :title="session.title"
        :description="session.description"
        :thumbnail="session.thumbnail"
        category="announcement"
        icon="code"
        :color="session.color"
        is-last
      />
    </div>
  </section>
</template>

<script>
export default {
  async asyncData({ $axios, store }) {
    store.commit('env/setCurrentNodeId', 0) // 设置当前所在node
    try {
      const [sessions] = await Promise.all([
        $axios.get('/api/apex/session/sessions'),
      ])
      console.log('sessions: ', sessions)
      return { sessions }
    } catch (e) {
      console.error(e)
    }
  },
  data() {
    return {
      apexSessions: [
        {
          date: new Date('2022-02-08'),
          title: '12-反抗',
          description: '赛季通行证皮肤：汗洛、新传奇疯玛吉、寻血猎犬威望皮肤',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2021-11-02'),
          title: '11-逃脱隐世',
          description: '',
          thumbnail: '/images/session/session-11.jpg',
          color: 'red',
        },
        {
          date: new Date('2021-08-03'),
          title: '10-羽化',
          description: '',
          thumbnail: '/images/session/session-10.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
        {
          date: new Date('2022-04-02'),
          title: 'aaaaaaaaaaaaa',
          description: 'aaaaaaaaaaaaaa',
          thumbnail: '/images/session/session-12.jpg',
          color: 'red',
        },
      ],
    }
  },
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
