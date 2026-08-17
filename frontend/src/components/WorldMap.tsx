import { useEffect, useRef } from "react";
import * as d3 from "d3";

import { feature } from "topojson-client";
import type { Place } from "../types/ingredients";
import { generateRouteSegments } from "../utils/route";

interface WorldMapProps {
  places: Place[];
  onPlaceClick?: (place: Place) => void;
}

export default function WorldMap({
  places,
  onPlaceClick,
}: WorldMapProps) {
  const svgRef = useRef<SVGSVGElement | null>(null);

  const width = 800;
  const height = 500;

  const projection = d3
    .geoNaturalEarth1()
    .scale(150)
    .translate([width / 2, height / 2]);

  const path = d3
    .geoPath()
    .projection(projection);

  // Draw the static world map once.
  useEffect(() => {
    async function drawMap() {
      if (!svgRef.current) return;

      const world = await fetch("/countries-110m.json")
        .then(res => res.json());

      const countries = feature(
        world, world.objects.countries
      );

      const svg = d3.select(svgRef.current);

      const countriesLayer = svg.select(".countries-layer");

      countriesLayer
        .selectAll("path.country")
        .data(countries.features)
        .join("path")
        .attr("class", "country")
        .attr("d", path)
        .attr("fill", "white")
        .attr("stroke", "gray");
    }

    drawMap();
  }, []);

  // Update places and routes.
  useEffect(() => {
    if (!svgRef.current) return;

    const svg = d3.select(svgRef.current);
    const placesLayer = svg.select(".places-layer");

    placesLayer
      .selectAll("circle.place")
      .data(places)
      .join(
        enter =>
          enter
            .append("circle")
            .attr("class", "place")
            .attr("fill", "red")
            .attr("r", 5)
            .style("cursor", "pointer"),

        update => update,
        exit => exit.remove()
      )
      .attr("cx", d =>
        projection([d.longitude, d.latitude])![0]
      )
      .attr("cy", d =>
        projection([d.longitude, d.latitude])![1]
      )
      .on("click", (event, place) => {
        onPlaceClick?.(place);
      });

    // Routes
    const routeSegments = generateRouteSegments(places);

    const routeLines = routeSegments.map(
      segment => [
        projection([
          segment.from.longitude,
          segment.from.latitude,
        ]),
        projection([
          segment.to.longitude,
          segment.to.latitude,
        ]),
      ]
    );

    // console.log("Route Lines:", routeLines);

    const routesLayer = svg.select(".routes-layer");

    routesLayer
      .selectAll("path.route")
      .data(routeLines)
      .join("path")
      .attr("class", "route")
      .attr("d", d3.line())
      .attr("fill", "none")
      .attr("stroke", "blue")
      .attr("stroke-width", 2);
  }, [places, onPlaceClick]);

  return (
    <svg
      ref={svgRef}
      width={width}
      height={height}
    >
      <g className="countries-layer" />
      <g className="routes-layer" />
      <g className="places-layer" />
    </svg>
  );
}