<template>
  <section class="main">
    <div class="container">
      <div class="main-body no-bg">
        <div class="widget signin">
          <div class="widget-header">注册</div>
          <div class="widget-content">
            <div class="field">
              <label class="label">昵称</label>
              <div class="control has-icons-left">
                <input
                  v-model="nickname"
                  class="input is-success"
                  type="text"
                  placeholder="请输入昵称"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-username" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">邮箱</label>
              <div class="control has-icons-left">
                <input
                  v-model="email"
                  class="input is-success"
                  type="text"
                  placeholder="请输入邮箱"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-email" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">密码</label>
              <div class="control has-icons-left">
                <input
                  v-model="password"
                  class="input"
                  type="password"
                  placeholder="请输入密码"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">确认密码</label>
              <div class="control has-icons-left">
                <input
                  v-model="rePassword"
                  class="input"
                  type="password"
                  placeholder="请再次输入密码"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">验证码</label>
              <div class="control has-icons-left">
                <div class="field is-horizontal">
                  <div class="field login-captcha-input">
                    <input
                      v-model="captchaCode"
                      class="input"
                      type="text"
                      placeholder="验证码"
                      @keyup.enter="signup"
                    />
                    <span class="icon is-small is-left"
                      ><i class="iconfont icon-captcha"
                    /></span>
                  </div>
                  <div v-if="captchaUrl" class="field login-captcha-img">
                    <a @click="showCaptcha"><img :src="captchaUrl" /></a>
                  </div>
                </div>
              </div>
            </div>

            <div class="field">
              <div class="control">
                <button class="button is-success" @click="signup">注册</button>
                <nuxt-link class="button is-text" to="/user/signin">
                  已有账号，前往登录&gt;&gt;
                </nuxt-link>
              </div>
            </div>

            <!-- <div
              v-if="loginMethod.qq || loginMethod.github || loginMethod.osc"
              class="third-party-line"
            >
              <div class="third-party-title">
                <span>第三方账号登录</span>
              </div>
            </div> -->

            <!-- <div class="third-parties">
              <github-login v-if="loginMethod.github" :ref-url="ref" />
              <osc-login v-if="loginMethod.osc" :ref-url="ref" />
              <qq-login v-if="loginMethod.qq" :ref-url="ref" />
            </div> -->
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  asyncData({ params, query }) {
    return {
      ref: query.ref,
    }
  },
  data() {
    return {
      nickname: '',
      email: '',
      password: '',
      rePassword: '',
      captchaId: '',
      captchaUrl: '',
      captchaCode: '',
    }
  },
  head() {
    return {
      title: this.$siteTitle('注册'),
    }
  },
  computed: {
    loginMethod() {
      return this.$store.state.config.config.loginMethod
    },
  },
  mounted() {
    this.showCaptcha()
  },
  methods: {
    async signup() {
      try {
        await this.$store.dispatch('user/signup', {
          captchaId: this.captchaId,
          captchaCode: this.captchaCode,
          nickname: this.nickname,
          email: this.email,
          password: this.password,
          rePassword: this.rePassword,
          ref: this.ref,
        })
        if (this.ref) {
          // 跳到登录前
          this.$linkTo(this.ref)
        } else {
          // 跳到个人主页
          this.$linkTo('/user/profile')
        }
      } catch (err) {
        this.$message.error(err.message || err)
        await this.showCaptcha()
      }
    },
    async showCaptcha() {
      try {
        const ret = await this.$axios.get('/api/captcha/request')
        this.captchaId = ret.captchaId
        this.captchaUrl =
          '/api/captcha/show?captchaId=' +
          this.captchaId +
          '&timestamp=' +
          new Date().getTime()
        this.captchaCode = ''
      } catch (e) {
        this.$message.error(e.message || e)
      }
    },
  },
}
</script>

<style lang="scss" scoped></style>
