<script setup>
import {reactive} from 'vue'
import {useVuelidate} from '@vuelidate/core'
import {email, required} from '@vuelidate/validators'
import Address from "@/components/form/Address.vue";

const initialState = {
  name: '',
  email: '',
  select: null,
  checkbox: null,
}

const state = reactive({
  ...initialState,
})

const items = [
  'Item 1',
  'Item 2',
  'Item 3',
  'Item 4',
]

const rules = {
  name: {required},
  email: {required, email},
  select: {required},
  items: {required},
  checkbox: {required},
}

const v$ = useVuelidate(rules, state)

function clear() {
  v$.value.$reset()

  for (const [key, value] of Object.entries(initialState)) {
    state[key] = value
  }
}

</script>

<template>
  <v-form fast-fail>
    <v-card title="Address" class="p-4">
      <Address/>
    </v-card>


    <v-card title="Info" class="p-4 mt-4">
      <v-text-field
          v-model="state.name"
          :counter="10"
          :error-messages="v$.name.$errors.map(e => e.$message)"
          label="Name"
          required
          @blur="v$.name.$touch"
          @input="v$.name.$touch"
      ></v-text-field>

      <v-text-field
          v-model="state.email"
          :error-messages="v$.email.$errors.map(e => e.$message)"
          label="E-mail"
          required
          @blur="v$.email.$touch"
          @input="v$.email.$touch"
      ></v-text-field>

      <v-select
          v-model="state.select"
          :error-messages="v$.select.$errors.map(e => e.$message)"
          :items="items"
          label="Item"
          required
          @blur="v$.select.$touch"
          @change="v$.select.$touch"
      ></v-select>
    </v-card>

    <v-checkbox
        v-model="state.checkbox"
        :error-messages="v$.checkbox.$errors.map(e => e.$message)"
        label="Do you agree?"
        required
        @blur="v$.checkbox.$touch"
        @change="v$.checkbox.$touch"
    ></v-checkbox>

    <v-btn
        class="me-4"
        @click="v$.$validate"
    >
      submit
    </v-btn>
    <v-btn @click="clear">
      clear
    </v-btn>
  </v-form>
</template>
<style scoped>

</style>