<!-- eslint-disable vue/require-v-for-key -->
<template id="postingan">
  <div>
      <h1>Postingan</h1>
      <div class="list-group">
          <a v-for="post in posts" href="#" class="list-group-item clearfix">
              {{ post }}
              <span class="pull-right">
                  <button class="btn btn-xs btn-info">
                      <span class="fa fa-edit"></span>
                  </button>
                  <button class="btn btn-xs btn-warning">
                      <span class="fa fa-trash"></span>
                  </button>
              </span>
          </a>
      </div>
  </div>
</template>

<script>
  export default {
    // eslint-disable-next-line vue/multi-word-component-names
    name: 'Postingan'
  }
  // eslint-disable-next-line vue/multi-word-component-names, no-undef
  Vue.component('postingan', {
    template: '#postingan',
    data: function() {
        return {
            posts: [
                'Membuat Dynamic Component',
                'Api Laravel 5.5',
                'HTML Fundamental',
                'Otentikasi Dengan Laravel 5.5'
             ]
        }
    }
  })
</script>