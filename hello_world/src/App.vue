<template>
	<div id="app" style="">
		<!-- <div id="app2"> -->
	<InputForm></InputForm>
		<!-- </div> -->
		
		<div v-for="affair in affairs" :key="affair.ID">
			<div class="row at-row">
				<div class="col-md-4">
					<div class="at-box-row bg-c-brand-dark common">{{affair.ID}}</div>
				</div>
				<div class="col-md-4">
					<div class="at-box-row bg-c-brand-light common">{{affair.events_title}}</div>
				</div>
				<div class="col-md-4">
					<div class="at-box-row bg-c-brand-dark common">{{affair.events_deadline}}</div>
				</div>
				<div class="col-md-4">
					<div class="at-box-row bg-c-brand-light common">{{affair.extre_sth}}</div>
				</div>
				<at-button v-on:click="deleteAffairs(affair.ID)" type="primary" style="float: left;">delete</at-button>
				<at-button v-on:click="modifyAffairs(affair.ID)" type="primary" style="float: left;">modify</at-button>
			</div>
		</div>
	</div>
</template>

<script>
	import axios from 'axios';
	import InputForm from "./components/InputForm";

	export default {
		name: 'app',
		components: {InputForm},
		data() {
			return {
				affairs: []
			}
		},
		methods: {
			deleteAffairs: function(ID) {
				// alert(ID);
				axios.delete('http://localhost:1221/opera/' + ID).then(() => {
					this.getAllAffairs()
				})
				
			},
			modifyAffairs: function(ID) {
				//回调
				alert(ID);
				this.getAllAffairs()
			},
			getAllAffairs: function() {
				axios.get('http://localhost:1221/allAffairs').then(res => {
					this.affairs = res.data.data
					// console.log(this.affairs)
				}).catch(err => {
					this.$Message.error(err)
				})
			}
		},
		mounted() {
			this.getAllAffairs();
		},
	}
</script>

<style>
	#app {
		font-family: 'Avenir', Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;
		color: #2c3e50;
		margin-top: 60px;
	}

	div.common {
		border-radius: 5px;
		background-color: #66B3F3;
		height: 30px;
	}
</style>
