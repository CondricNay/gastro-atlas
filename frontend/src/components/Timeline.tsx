import { useEffect, useRef } from "react";
import * as d3 from "d3";

interface TimelineMarker {
  year: number;
  label: string;
}

interface TimelineProps {
  minYear: number;
  maxYear: number;
  currentYear: number;
  onYearChange: (year: number) => void;
  markers: TimelineMarker[];
}

export default function Timeline({
  minYear,
  maxYear,
  currentYear,
  onYearChange,
  markers,
}: TimelineProps) {
    const svgRef = useRef<SVGSVGElement | null>(null);

    const WIDTH = 800;
    const HEIGHT = 120;

    const margin = {
        left: 40,
        right: 40,
    };

    const xScale = d3
        .scaleLinear()
        .domain([minYear, maxYear])
        .range([margin.left, WIDTH - margin.right])
        .clamp(true);

    useEffect(() => {
        if (!svgRef.current) return;

        const svg = d3.select(svgRef.current);
        svg.selectAll("*").remove();
        
        const track = svg.append("rect")
            .attr("x", margin.left)
            .attr("y", HEIGHT / 2 - 4)
            .attr("width", WIDTH - margin.left - margin.right)
            .attr("height", 8)
            .attr("rx", 4);

        track.on("click", (event) => {
            const [mouseX] = d3.pointer(event);

            const year = Math.round(
                xScale.invert(mouseX)
            );

            onYearChange(year);
        });

        const handle = svg.append("circle")
            .attr("cx", xScale(currentYear))
            .attr("cy", HEIGHT / 2)
            .attr("r", 10)
            .attr("fill", "#2563eb");

        const drag = d3.drag<SVGCircleElement, unknown>()
            .on("drag", (event) => {
                const year = Math.round(
                    xScale.invert(event.x)
                );

                onYearChange(year);
            });

        handle.call(drag);

        svg.append("g")
            .selectAll("circle")
            .data(markers)
            .join("circle")
            .attr("cx", d => xScale(d.year))
            .attr("cy", HEIGHT / 2)
            .attr("r", 4)
            .attr("fill", "#6b7280");

        svg.append("g")
            .selectAll("text")
            .data(markers)
            .join("text")
            .attr("x", d => xScale(d.year))
            .attr("y", HEIGHT / 2 + 25)
            .attr("text-anchor", "middle")
            .text(d => d.year);
            

        svg.append("text")
            .attr("x", xScale(currentYear))
            .attr("y", 20)
            .attr("text-anchor", "middle")
            .text(currentYear);

    }, [currentYear, markers, minYear, maxYear, onYearChange]);

    return <svg
        ref={svgRef}
        width={WIDTH}
        height={HEIGHT}
    />
}
