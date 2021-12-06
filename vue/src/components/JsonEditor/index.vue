<template>
  <div class="json-editor">
    <textarea ref="textarea" />
  </div>
</template>

<script>
import CodeMirror from 'codemirror'
import 'codemirror/addon/lint/lint.css'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/rubyblue.css'
require('script-loader!jsonlint')
import 'codemirror/mode/javascript/javascript'
import 'codemirror/addon/lint/lint'
import 'codemirror/addon/lint/json-lint'

export default {
  name: 'JsonEditor',
  /* eslint-disable vue/require-prop-types */
  props: ['value'],
  data() {
    return {
      jsonEditor: false
    }
  },
  watch: {
    value(value) {
      const editorValue = this.jsonEditor.getValue()
      if (value !== editorValue) {
        //var tempJson = JSON.parse(this.value)
        //this.jsonEditor.setValue(JSON.stringify(tempJson, null, '\t'), null, 2)
        this.jsonEditor.setValue(this.value)
      }
    }
  },
  mounted() {
    this.jsonEditor = CodeMirror.fromTextArea(this.$refs.textarea, {
      lineNumbers: true,
      mode: 'application/json',
      gutters: ['CodeMirror-lint-markers'],
      theme: 'rubyblue',
      lint: true,
      extraKeys: {
        'F7': function autoFormat(editor) {
          var tempJson = JSON.parse(editor.getValue())
          editor.setValue(JSON.stringify(tempJson, null, '\t'))
        }// 代码格式化
      }
    })
    if (this.value != '') {
      var tempJson = JSON.parse(this.value)
      this.jsonEditor.setValue(JSON.stringify(tempJson, null, '\t'))
    }
    this.jsonEditor.on('change', cm => {
      this.$emit('changed', cm.getValue())
      this.$emit('input', cm.getValue())
    })
  },
  methods: {
    getValue() {
      console.log(this.jsonEditor.getValue(), 'this.jsonEditor.getValue()')
      return this.jsonEditor.getValue()
    }
  },
  destroyed() {
    console.log("啊 我死了")
  }
}
</script>

<style scoped>
.json-editor{
  height: 100%;
  position: relative;
}
.json-editor >>> .CodeMirror {
  height: auto;
  min-height: 300px;
}
.json-editor >>> .CodeMirror-scroll{
  min-height: 300px;
}
.json-editor >>> .cm-s-rubyblue span.cm-string {
  color: #F08047;
}
</style>
