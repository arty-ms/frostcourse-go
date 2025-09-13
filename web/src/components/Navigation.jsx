export default function Navigation({ currentStep, totalSteps, onPrev, onNext, canGoNext }) {
    return (
        <div className="flex justify-between items-center mt-6 pt-6 border-t border-tg-section-separator">
            <button
                className={`
          bg-tg-secondary-bg text-tg-text border border-tg-section-separator
          px-6 py-3 rounded-lg font-medium transition-all duration-200
          ${currentStep === 1 ? 'opacity-50 cursor-not-allowed' : 'hover:bg-tg-bg'}
        `}
                onClick={onPrev}
                disabled={currentStep === 1}
            >
                ← Назад
            </button>

            <span className="text-sm text-tg-hint font-medium">
        {currentStep} из {totalSteps}
      </span>

            <button
                className={`
          bg-tg-button text-tg-button-text px-6 py-3 rounded-lg font-medium 
          transition-all duration-200
          ${(!canGoNext || currentStep === totalSteps) ? 'opacity-50 cursor-not-allowed' : 'hover:opacity-90'}
        `}
                onClick={onNext}
                disabled={!canGoNext || currentStep === totalSteps}
            >
                Далее →
            </button>
        </div>
    );
}