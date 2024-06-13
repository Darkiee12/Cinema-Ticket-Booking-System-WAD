import { useState } from "react";
import { FaArrowLeft, FaArrowRight } from "react-icons/fa";

type ImageSliderProps = {
    imageUrls: string[];
}
export function ImageSlider({imageUrls}: ImageSliderProps) {
    const [currentImageIndex, setCurrentImageIndex] = useState(0)

    function showNextImage() {
        if(currentImageIndex === imageUrls.length - 1) {
            setCurrentImageIndex(0)
        } else {
            setCurrentImageIndex(currentImageIndex + 1)
        }
    }

    function showPreviousImage() {
        if(currentImageIndex === 0) {
            setCurrentImageIndex(imageUrls.length - 1)
        } else {
            setCurrentImageIndex(currentImageIndex - 1)
        }
    }
    return (
        <div className="w-full h-full relative">
            <img src = {imageUrls[currentImageIndex]} className="h-full w-full block object-cover" />
            <button onClick = {showPreviousImage} className="block absolute top-0 bottom-0 p-1 hover:bg-black hover:bg-opacity-50"><FaArrowLeft className="fill-white w-6 h-6" /></button>
            <button onClick = {showNextImage} className="block absolute top-0 bottom-0 p-1 right-0 hover:bg-black hover:bg-opacity-50"><FaArrowRight className="fill-white w-6 h-6"/></button>
        </div>
    );      
};
