<?php

namespace App\Http\Requests;

class ClassSubjectMaterialRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'class_subject_id' => 'required|numeric|gt:0',
            'description' => 'required|string|max:500',
            'content' => 'required|string|max:2500',
            'date' => 'required|date',
        ];
    }

    /**
     * @return array
     * Custom validation message
     */
    public function messages(): array
    {
        return [
            'class_subject_id.required' => 'The class subject ID field is required.',
            'class_subject_id.numeric' => 'The class subject ID must be a number.',
            'class_subject_id.gt' => 'The class subject ID must be a positive number.',
            'description.required' => 'The description field is required.',
            'description.string' => 'The description must be a string.',
            'description.max' => 'The description may not be greater than :max characters.',
            'content.required' => 'The content field is required.',
            'content.string' => 'The content must be a string.',
            'content.max' => 'The content may not be greater than :max characters.',
            'date.required' => 'The date field is required.',
            'date.date' => 'The date must be a valid date.',
        ];
    }
}
