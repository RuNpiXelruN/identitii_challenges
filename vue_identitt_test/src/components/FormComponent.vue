<template>
    <transition name="fade" appear>
        <div class="form-component">
            <v-layout justify-center>            
                <v-flex xs12 sm6 class="form-flex-wrapper">
                    <v-layout row justify-start>
                        <router-link tag="v-btn" class="blue accent-3 back-btn v-btn--outline v-btn--depressed blue--text text--accent-3" to="/">Home</router-link>
                    </v-layout>
                    <h1 class="form-component-h1">Identitii - Attribute Picker Form</h1>
                    <div class="form-info-wrapper">
                        <h3 class="form-heading">YOUR SHARED DETAILS</h3>
                        <p class="form-info">For optimal speed of settlement, the following details about you will be shared with parties in the transaction.</p>
                    </div>
                    <v-layout row justify-start>
                    <v-form ref="attributeForm" class="attribute-form">
                        <p class="form-incl">Incl.</p>
                        <v-layout row justify-space-between>
                            <v-text-field
                                id="sender/name"
                                ref="nameField"
                                :rules="nameRules"
                                class="identitii-form-input"
                                label="*Name"
                                v-model="name"
                                required>
                            </v-text-field>
                            <v-checkbox
                                class="form-checkbox-right"
                                v-model="nameCheckbox"
                                disabled
                                light>
                            </v-checkbox>
                        </v-layout>
                        <v-layout row justify-space-between>
                            <v-select
                                id="sender/entityType"
                                ref="entityField"
                                class="identitii-form-input"
                                label="Entity Type"
                                v-model="selectedEntityType"
                                :items="entities">
                            </v-select>
                            <v-checkbox
                                class="form-checkbox-right"
                                :disabled="selectedEntityType.length == 0"
                                v-model="entityCheckbox"
                                light>
                            </v-checkbox>
                        </v-layout>
                        <v-layout row justify-space-between>
                            <v-select
                                id="sender/country"
                                ref="countryField"
                                :rules="countryRules"
                                class="identitii-form-input"
                                label="*Country"
                                v-model="selectedCountry"
                                :items="countries">
                            </v-select>
                            <v-checkbox
                                class="form-checkbox-right"
                                v-model="countryCheckbox"
                                disabled
                                light>
                            </v-checkbox>
                        </v-layout>
                        <v-layout row justify-space-between>
                            <v-textarea
                                id="sender/address"
                                ref="addressField"
                                class="identitii-form-input"
                                label="Address"
                                v-model="address"
                            >                            
                            </v-textarea>                        
                            <v-checkbox
                                v-model="addressCheckbox"
                                :disabled="!address"
                                class="form-checkbox-right"
                                light>
                            </v-checkbox>
                        </v-layout>
                        <v-layout row justify-space-between>
                            <v-text-field
                                id="sender/postcode"
                                ref="postCodeField"
                                class="identitii-form-input"
                                label="Post Code"
                                v-model="postCode">                            
                            </v-text-field>
                            <v-checkbox
                                v-model="postCodeCheckbox"
                                class="form-checkbox-right"
                                :disabled="!postCode"
                                light>
                            </v-checkbox>
                        </v-layout>
                        <v-layout row justify-space-between>
                            <v-text-field
                                id="sender/email"
                                ref="emailField"
                                class="identitii-form-input"
                                label="Email"
                                :rules="emailRules"
                                @blur="handleBlur"
                                @keyup="checkErrState"
                                v-model="email">                            
                            </v-text-field>
                            <v-checkbox
                                v-model="emailCheckbox"
                                class="form-checkbox-right"
                                :disabled="!emailValid"
                                light>
                            </v-checkbox>
                        </v-layout>                        

                        <v-layout row justify-space-between>
                            <v-btn
                                outline
                                class="reset-btn"
                                color="pink lighten-1"
                                @click="resetForm"
                                >Reset                            
                            </v-btn>
                            <v-btn
                                color="blue accent-1 white--text"
                                :disabled="!valid"
                                @click.prevent="handleSubmit"
                                >Submit                            
                            </v-btn>
                        </v-layout>
                    </v-form>
                    <attribute-output :result="result" />
                    </v-layout>
                </v-flex>
            </v-layout>
        </div>        
    </transition>
</template>
<script>
import AttributeOutput from '@/components/AttributeOutput.vue'
export default {
    components: {
        AttributeOutput,
    },
    data() {
        return {
            name: null,
            nameCheckbox: true,
            nameRules: [
                (v) => !!v || 'Name is required',
            ],
            selectedEntityType: [],
            entityCheckbox: false,
            entities: [],
            selectedCountry: [],
            countryCheckbox: true,
            countryRules: [
                (v) => !!v || "Country is required"
            ],
            countries: [],
            address: null,
            addressCheckbox: false,
            postCode: null,
            postCodeCheckbox: false,
            email: null,
            emailCheckbox: false,
            emailValid: false,
            emailRules: [
                (v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
            ],
            formData: null,
            result: null,
        }
    },
    watch: {
        formData(newVal, oldVal) {
            if (newVal) {
                this.setData()
            }            
        },
    },
    computed: {
        valid() {            
            return (this.name && !!this.selectedCountry.length)
        },
        emailValidations() {            
            return this.$refs.emailField
        },
    },
    created() {
        this.formData = this.$store.getters.formData
    },
    methods: {        
        checkErrState() {
            let errState = this.$refs.emailField.$options._parentVnode.componentInstance.validationState
            if ((this.email && this.email.length > 0) && errState == "primary") {
                this.emailValid = true
            } else {
                this.emailValid = false
                this.emailCheckbox = false
            }
        },        
        handleBlur() {
            let errState = this.$refs.emailField.$options._parentVnode.componentInstance.validationState
            
            if (this.email == null || this.email.length < 1) {
                this.emailValid = false
                this.emailCheckbox = false
                this.$refs.emailField.reset()
            }
        },
        setData() {
            this.entities = this.formData.formTypes.entityType.uiOptions.options.map(type => {
                return {
                    value: type.value,
                    text: type.name
                }
            })

            this.countries = this.formData.formTypes.country.uiOptions.options.map(type => {
                return {
                    value: type.value,
                    text: type.name
                }
            })

            this.setInitialData()
        },

        handleSubmit() {
            this.emailValidations.$options._parentVnode.componentInstance.validationState
            let data = {selected: []}            
            data.selected.push({id: this.$refs.nameField.$attrs.id, contentType: "text/string", value: this.name})
            data.selected.push({id: this.$refs.countryField.$attrs.id, contentType: "text/string", value: this.selectedCountry})

            if (this.entityCheckbox) {
                data.selected.push({id: this.$refs.entityField.$attrs.id, contentType: "text/string", value: this.selectedEntityType})
            }
            
            if (this.addressCheckbox) {
                data.selected.push({id: this.$refs.addressField.$attrs.id, contentType: "text/string", value: this.address})
            }
            
            if (this.postCodeCheckbox) {
                data.selected.push({id: this.$refs.postCodeField.$attrs.id, contentType: "", value: this.postCode})
            }
            
            if (this.emailCheckbox) {
                data.selected.push({id: this.$refs.emailField.$attrs.id, contentType: "text/string", value: this.email})
            }
            
            this.result = JSON.stringify(data)
            this.resetForm()
            this.emailValid = false
        },

        setInitialData() {
            let data = {
                name: null,
                nameCheckbox: true,                
                selectedEntityType: [],
                entityCheckbox: false,
                selectedCountry: [],
                countryCheckbox: true,
                address: null,
                addressCheckbox: false,
                postCode: null,
                postCodeCheckbox: false,
                email: null,
                emailCheckbox: false,
            }            

            return data
        },

        resetForm() {
            let initData = this.setInitialData()
            for (let prop in initData) {
                this[prop] = initData[prop]
            }
            this.$refs.emailField.reset()
            this.$refs.nameField.reset()
        }    
    }
}
</script>
<style lang="scss">
    .form-component {
        margin-top: 60px;

        .back-btn {
            margin-bottom: 30px;
        }

        .form-component-h1 {
            margin-bottom: 30px;
            text-align: left;
        }
    }
    .form-info-wrapper {
        display: flex;
        flex-flow: column nowrap;
        align-items: flex-start;
        width: 100%;

        .form-heading {
            width: 100%;
            text-align: left;
            border-bottom: 1px solid grey;
        }
        .form-info {
            text-align: left;
            margin-bottom: 60px;
        }        
    }

    .attribute-form {
        width: 60%;

        .form-incl {
            margin-bottom: -10px;
            text-align: right;
            margin-right: 5px;
        }
    }

    .identitii-form-input {
        width: 60%;
    }

    .form-checkbox-right {
        .v-input__slot {
            justify-content: flex-end;
        }
    }

    .reset-btn {
        margin-left: 0 !important;
    }

    // animation / transition
    .fade-enter {
        opacity: 0
    }

    .fade-enter-active {
        transition: opacity 1s ease;
    }
    
    // .fade-leave {}

    .fade-leave-active {
        opacity: 0;
    }
</style>


