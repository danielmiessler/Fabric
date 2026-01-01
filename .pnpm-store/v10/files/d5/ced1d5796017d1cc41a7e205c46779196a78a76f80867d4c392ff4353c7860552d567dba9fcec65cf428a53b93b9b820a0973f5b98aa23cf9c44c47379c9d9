<!-- NoirLight: `filter: url(#NoirLight)` -->
<svg id="svg-filter-noirlight" class="filter absolute -left-full w-0 h-0">
	<filter
		id="NoirLight"
		x="-20%"
		y="-20%"
		width="140%"
		height="140%"
		filterUnits="objectBoundingBox"
		primitiveUnits="userSpaceOnUse"
		color-interpolation-filters="linearRGB"
	>
		<feColorMatrix type="saturate" values="0" x="0%" y="0%" width="100%" height="100%" in="SourceGraphic" result="colormatrix2" />
		<feBlend mode="saturation" x="0%" y="0%" width="100%" height="100%" in="SourceGraphic" in2="colormatrix2" result="blend2" />
		<feBlend mode="screen" x="0%" y="0%" width="100%" height="100%" in="colormatrix2" in2="blend2" result="blend3" />
		<feColorMatrix type="luminanceToAlpha" x="0%" y="0%" width="100%" height="100%" in="blend3" result="colormatrix3" />
		<feBlend mode="exclusion" x="0%" y="0%" width="100%" height="100%" in="blend3" in2="colormatrix3" result="blend5" />
	</filter>
</svg>
