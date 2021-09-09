<template>
  <div>
    <b-select expanded v-bind:value="value" @input="$emit('input', parseInt($event))" placeholder="Answer type...">
      <option value="1">Number</option>
      <option value="2">Rating</option>
      <option value="3">Weekday</option>
    </b-select>
  </div>
</template>

<script>
export default {
  name: "FormFieldProfileQuestionAnswerType",

  props: {
    value: {
      type: Number,
    }
  },


}
</script>

<style scoped>

</style>