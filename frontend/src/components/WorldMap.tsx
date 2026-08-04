import { useEffect, useRef } from "react";
import * as d3 from "d3";

import { feature } from "topojson-client";
import type { Place } from "../types/ingredients";

interface WorldMapProps {
  places: Place[];
}

export default function WorldMap({ places }: WorldMapProps) {
  const svgRef = useRef<SVGSVGElement | null>(null);

  const width = 800;
  const height = 500;

  const projection = d3.geoNaturalEarth1()
    .scale(150)
    .translate([width / 2, height / 2]);

  const path = d3.geoPath()
    .projection(projection);


  useEffect(() => {
    async function drawMap() {
      const world = await fetch("/countries-110m.json")
        .then(res => res.json());

      // Convert TopoJSON to GeoJSON
      const countries = feature(world, world.objects.countries);

      const svg = d3.select(svgRef.current);

      svg.selectAll("path")
        .data(countries.features)
        .join("path")
        .attr("d", path)
        .attr("fill", "white")
        .attr("stroke", "gray");

      // Circles for places
      svg.selectAll("circle")
      .data(places)
      .join(
        enter => enter
            .append("circle")
            .attr("fill", "red")
            .attr("r", 5),
        update => update,
        exit => exit.remove()
      )
      .attr("cx", d =>
        projection([d.longitude, d.latitude])![0]
      )
      .attr("cy", d =>
        projection([d.longitude, d.latitude])![1]
      );

      const routePoints = places
        .map(place =>
          projection([
            place.longitude,
            place.latitude
          ])
        );

      svg.selectAll("path.route")
        .data([routePoints])
        .join("path")
        .attr("class", "route")
        .attr("d", d3.line())
        .attr("fill", "none")
        .attr("stroke", "blue")
        .attr("stroke-width", 2);
    }

    drawMap();
  }, [places]);

  return (
    <svg
      ref={svgRef}
      width={width}
      height={height}
    />
  );
}