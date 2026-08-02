import { useEffect, useRef } from "react";
import * as d3 from "d3";

interface Place {
  name: string;
  relationship: string;
  startYear: number;
  endYear: number | null;
}

interface TimelineProps {
  places: Place[];
}


export default function Timeline({ places }: TimelineProps) {
  const svgRef = useRef<SVGSVGElement | null>(null);
  const CURRENT_YEAR = new Date().getFullYear();

  useEffect(() => {
    if (!svgRef.current) return;

    const width = 800;
    const height = places.length * 60 + 80;

    const svg = d3.select(svgRef.current);

    svg.selectAll("*").remove();

    const margin = {
      top: 30,
      right: 30,
      bottom: 40,
      left: 150
    };

    const years = places.flatMap(p => [
      p.startYear, p.endYear ?? new Date().getFullYear()
    ]);

    const x = d3.scaleLinear()
      .domain([d3.min(years)!, d3.max(years)!])
      .range([
        margin.left,
        width - margin.right
      ]);


    const y = d3.scaleBand()
      .domain(places.map(p => p.name))
      .range([
        margin.top,
        height - margin.bottom
      ])
      .padding(0.3);


    // Axis
    svg.append("g")
      .attr(
        "transform",
        `translate(0,${height - margin.bottom})`
      )
      .call(
        d3.axisBottom(x)
      );


    // Labels
    svg.append("g")
      .selectAll("text")
      .data(places)
      .join("text")
      .attr("x", margin.left - 10)
      .attr("y", p =>
        y(p.name)! + y.bandwidth()/2
      )
      .attr("text-anchor", "end")
      .attr("alignment-baseline", "middle")
      .text(p => p.name);


    // Timeline bars
    svg.append("g")
      .selectAll("rect")
      .data(places)
      .join("rect")
      .attr("x", p => x(p.startYear))
      .attr("y", p => y(p.name)!)
      .attr("width", p =>
        x(p.endYear ?? CURRENT_YEAR) - x(p.startYear)
      )
      .attr("height", y.bandwidth());


  }, [places]);


  return (
    <svg
      ref={svgRef}
      width="800"
      height={places.length * 60 + 80}
    />
  );
}